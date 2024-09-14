package sms

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOTP(t *testing.T) {
	data := NewSMSDevelopment()
	err := data.SendOTP("9395756899", 1000)

	require.NoError(t, err)
}
