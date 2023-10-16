package adapters

import (
	"context"
	"ecommerce/internal/entities"
	"fmt"
	"net/smtp"

	"github.com/go-redis/redis/v8"
)

type OtpAdapter struct {
	redisClient *redis.Client
	smtpConfig  SMTPConfig
}

type SMTPConfig struct {
	SMTPServer   string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
	Receiver     string
}

func NewOtpAdapter(redisclient *redis.Client, smtpconfig SMTPConfig) *OtpAdapter {
	return &OtpAdapter{
		redisClient: redisclient,
		smtpConfig:  smtpconfig,
	}
}

func (otpadapter *OtpAdapter) SaveOTP(otp entities.OTP) error {

	ctx := context.Background()

	return otpadapter.redisClient.Set(ctx, otp.Email, otp.Value, otp.ExpiresAt).Err()
}

func (otpadapter *OtpAdapter) GetOTP(email string) (string, error) {

	ctx := context.Background()

	otp, err := otpadapter.redisClient.Get(ctx, email).Result()

	if err != nil {
		return "", err
	}

	return otp, nil

}

func (otpadapter *OtpAdapter) SendOTP(otp string) error {

	auth := smtp.PlainAuth("", otpadapter.smtpConfig.SMTPUsername, otpadapter.smtpConfig.SMTPPassword, otpadapter.smtpConfig.SMTPServer)

	message := "To: " + otpadapter.smtpConfig.Receiver + "\r\n" +
		"Subject: OTP Verification\r\n" +
		"\r\n" +
		"Here is your otp,\r\n" + otp +
		"\r\n"

	if err := smtp.SendMail(otpadapter.smtpConfig.SMTPServer+":"+otpadapter.smtpConfig.SMTPPort, auth, otpadapter.smtpConfig.SMTPUsername, []string{otpadapter.smtpConfig.Receiver}, []byte(message)); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
