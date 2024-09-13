package otp_manager

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

const PhoneOtpTTL = time.Minute * 2

type OtpManager struct {
	redisClient *redis.Client
}

func NewOtpManger(redisClient *redis.Client) *OtpManager {
	return &OtpManager{redisClient}
}

func (m *OtpManager) Generate(ctx context.Context, phoneNumber string) (int, error) {
	otpCode := GenerateOtp(6)
	cmd := m.redisClient.Set(ctx, fmt.Sprintf("OTP:%v", phoneNumber), otpCode, PhoneOtpTTL)
	if cmd.Err() != nil {
		return 0, cmd.Err()
	}

	return otpCode, nil
}

func (m *OtpManager) Compare(ctx context.Context, phoneNumber string, otpCode int) (bool, error) {
	cmd := m.redisClient.Get(ctx, fmt.Sprintf("OTP:%v", phoneNumber))
	if cmd.Err() != nil {
		return false, cmd.Err()
	}

	otpStr, err := cmd.Result()
	if err != nil {
		return false, err
	}

	localOtpCode, err := strconv.Atoi(otpStr)
	if err != nil {
		return false, err
	}

	if localOtpCode == otpCode {
		return true, nil
	}

	return false, nil
}

func (m *OtpManager) Exist(ctx context.Context, phoneNumber string) bool {
	cmd := m.redisClient.Get(ctx, fmt.Sprintf("OTP:%v", phoneNumber))
	if cmd.Err() != nil {
		return false
	}

	otpStr, err := cmd.Result()
	if err != nil || len(otpStr) == 0 {
		return false
	}

	return true
}

func (m *OtpManager) Remove(ctx context.Context, phoneNumber string) error {
	cmd := m.redisClient.Del(ctx, fmt.Sprintf("OTP:%v", phoneNumber))
	if cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}
