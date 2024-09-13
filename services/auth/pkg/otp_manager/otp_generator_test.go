package otp_manager

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateOTP(t *testing.T) {
	otp := GenerateOtp(4)
	require.NotZero(t, otp)
}
