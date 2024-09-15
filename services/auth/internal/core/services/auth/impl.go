package auth_service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/geekengineers/Microservice-Project-Demo/services/auth/internal/core/domain/user"
	"github.com/geekengineers/Microservice-Project-Demo/services/auth/internal/ports"
	auth_manager "github.com/tahadostifam/go-auth-manager"

	"github.com/geekengineers/Microservice-Project-Demo/services/auth/pkg/otp_manager"
	"github.com/geekengineers/Microservice-Project-Demo/services/auth/pkg/sms"
)

const AccessTokenTTL = 20 * 24 * time.Hour

type Requirements struct {
	OtpManager  *otp_manager.OtpManager
	AuthManager auth_manager.AuthManager
	Repo        ports.AuthRepositorySecondaryPort
	SmsService  sms.Service
}

type Service struct {
	requirements *Requirements
}

func NewService(requirements *Requirements) *Service {
	return &Service{requirements}
}

func (s *Service) Login(ctx context.Context, phoneNumber string) (int, error) {
	_, err := s.requirements.Repo.FindByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		// User does not exists and let's create it
		newUser := &user.User{
			PhoneNumber: phoneNumber,
		}

		_, err := s.requirements.Repo.Create(ctx, newUser)
		if err != nil {
			return 0, ErrCreation
		}
	}

	otp, err := s.requirements.OtpManager.Generate(ctx, phoneNumber)
	if err != nil {
		return 0, ErrOtpCodeGeneration
	}

	// TODO - Make queue for sending sms later
	err = s.requirements.SmsService.SendOTP(phoneNumber, otp)
	if err != nil {
		return 0, ErrSendingSms
	}

	return otp, nil
}

func (s *Service) SubmitOtp(ctx context.Context, phoneNumber string, otpCode int) (*user.User, string, error) {
	u, err := s.requirements.Repo.FindByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		return nil, "", ErrPermissionDenied
	}

	valid, err := s.requirements.OtpManager.Compare(ctx, phoneNumber, otpCode)
	if err != nil || !valid {
		return nil, "", ErrPermissionDenied
	}

	accessToken, err := s.requirements.AuthManager.GenerateAccessToken(ctx, fmt.Sprintf("%d", u.ID), AccessTokenTTL)
	if err != nil {
		return nil, "", ErrTokenGeneration
	}

	return u, accessToken, nil
}

func (s *Service) Authenticate(ctx context.Context, token string) (*user.User, error) {
	claims, err := s.requirements.AuthManager.DecodeAccessToken(ctx, token)
	if err != nil || claims.Payload.UUID == "" {
		return nil, ErrPermissionDenied
	}

	if id, err := strconv.ParseInt(claims.Payload.UUID, 10, 64); err == nil {
		u, err := s.requirements.Repo.Find(ctx, id)
		if err != nil {
			return nil, ErrPermissionDenied
		}

		return u, nil
	}

	return nil, ErrPermissionDenied
}
