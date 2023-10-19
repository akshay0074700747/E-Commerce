package middlewares

import (
	"context"
	"ecommerce/internal/adapters"
	"ecommerce/internal/usecases"

	"ecommerce/web/config"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
)

var otpadapter *adapters.OtpAdapter

func Otp_Gen(email string) error {

	if emailval := strings.HasSuffix(email, "@gmail.com"); !emailval {
		return fmt.Errorf("this is not a valid email")
	}

	config, err := config.LoadConfig()

	if err != nil {
		panic("Cannot load env files...")
	}

	smtp_username := config.EMAIL
	smtp_password := config.PASSWORD

	// smtp_username := "akshay8547104@gmail.com"
	// smtp_password := "fqih bpdj fxrr obdp"

	smtp_server := "smtp.gmail.com"
	smtp_port := "587"

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	pong, err := client.Ping(ctx).Result()

	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return err
	}

	fmt.Println(pong)

	smtpconfig := adapters.SMTPConfig{
		SMTPServer:   smtp_server,
		SMTPPort:     smtp_port,
		SMTPUsername: smtp_username,
		SMTPPassword: smtp_password,
		Receiver:     email,
	}

	otpadapter = adapters.NewOtpAdapter(client, smtpconfig)

	otp := usecases.NewOtpUseCase(otpadapter)

	otpstring, err := otp.GenerateOtp(email)

	if err := otp.OTPRepository.SendOTP(otpstring); err != nil {
		fmt.Println(err.Error())
		return err
	}

	if err != nil {
		fmt.Println("Error generating otp :", err)
		return err
	}

	fmt.Println(otpstring)

	return nil
}

func Otp_Verify(email, otp string) (error, bool) {

	if emailval := strings.HasSuffix(email, "@gmail.com"); !emailval {
		return fmt.Errorf("this is not a valid email"), false
	}

	fmt.Println(otp)
	fmt.Println("this issssssssssssssssssssssssssssss")

	otpstring, err := otpadapter.GetOTP(email)

	if err != nil {
		fmt.Println(err.Error())
		return err, false
	}

	fmt.Println("this is the retrieved otp")
	fmt.Println(otpstring)

	if otp != otpstring {
		fmt.Println("otp doesnt match....")
		return fmt.Errorf("otp doesnt match"), false
	}

	return nil, true
}
