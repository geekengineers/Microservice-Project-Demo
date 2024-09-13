package sms

import "log"

type devel_service struct{}

func NewSMSDevelopment() Service {
	return &devel_service{}
}

func (s *devel_service) SendOTP(phoneNumber string, code int) error {
	log.Println("SMS SENT => ")
	log.Println("   === Phone Number: ", phoneNumber)
	log.Println("   === OTP Code: ", code)

	return nil
}
