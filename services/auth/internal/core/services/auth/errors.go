package auth_service

import "errors"

var (
	ErrCreation          = errors.New("user creation failed")
	ErrOtpCodeGeneration = errors.New("failed to generate otp code")
	ErrSendingSms        = errors.New("failed send sms")
	ErrPermissionDenied  = errors.New("permission denied")
	ErrTokenGeneration   = errors.New("failed to generate token")
)
