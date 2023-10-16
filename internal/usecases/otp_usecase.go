package usecases

import (
	"ecommerce/internal/entities"
	"ecommerce/internal/interfaces/repositories"
	"fmt"
	"math/rand"
	"time"
)

type OtpUseCase struct {
	OTPRepository repositories.OtpRepo
}

func NewOtpUseCase(OTPRepository repositories.OtpRepo) *OtpUseCase {

	return &OtpUseCase{
		OTPRepository: OTPRepository,
	}

}

func (OtpUseCase *OtpUseCase) GenerateOtp(email string) (string, error) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	otpstring := fmt.Sprintf("%06d", r.Intn(1000000))

	duration := 5 * time.Minute

	otp := entities.OTP{
		Value:     otpstring,
		Email:     email,
		ExpiresAt: duration,
	}

	if err := OtpUseCase.OTPRepository.SaveOTP(otp); err != nil {
		return "", err
	}

	return otpstring, nil

}

func (OtpUseCase *OtpUseCase) VerifyOtp(email, otp string) bool {

	recotp, err := OtpUseCase.OTPRepository.GetOTP(email)

	if err != nil || otp != recotp {
		fmt.Println(err.Error())
		return false
	}

	return true

}
