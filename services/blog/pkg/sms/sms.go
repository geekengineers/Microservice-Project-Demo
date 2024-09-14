package sms

type Service interface {
	SendOTP(phoneNumber string, code int) error
}

type service struct{}

func NewSmsService(phoneNumber string) Service {
	return &service{}
}

func (s *service) SendOTP(phoneNumber string, code int) error {
	panic("unimplemented")
}
