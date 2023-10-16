package repositories

import "ecommerce/internal/entities"

type OtpRepo interface {
	SaveOTP(otp entities.OTP) error
	GetOTP(email string) (string, error)
	SendOTP(otp string) error
}