package otp_manager

import (
	"crypto/rand"
	"math"
	"math/big"
	"strconv"
)

func GenerateOtp(length int) int {
	min := int64(math.Pow(10, float64(length)-1))
	max := int64(math.Pow(10, float64(length))) - 1

	randomNumber, err := rand.Int(rand.Reader, big.NewInt(max-min))
	if err != nil {
		panic(err)
	}

	number := int(randomNumber.Int64()) + int(min)

	if len(strconv.Itoa(number)) != length {
		return 0
	}

	return number
}
