package middlewares

import (
	"context"
	"ecommerce/internal/adapters"
	"ecommerce/internal/usecases"
	"ecommerce/web/config"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var otpadapter *adapters.OtpAdapter

func Otp_Gen(c *gin.Context) (error, bool) {

	config, err := config.LoadConfig()

	if err != nil {
		panic("Cannot load env files...")
	}

	smtp_username := config.EMAIL
	smtp_password := config.PASSWORD

	smtp_server := "smtp.gmail.com"
	smtp_port := "587"

	var jsondta map[string]string

	if err := c.Bind(&jsondta); err != nil {
		return err, false
	}

	if jsondta["otp"] != "" {

		fmt.Println("going to next")
		if err := Otp_Verify(jsondta["email"], jsondta["otp"]); err != nil {
			return err, false
		}

		fmt.Println("comin from the next")

		return nil, true

	} else {

		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

		ctx := context.Background()

		pong, err := client.Ping(ctx).Result()

		if err != nil {
			fmt.Println("Error connecting to Redis:", err)
			return err, false
		}

		fmt.Println(pong)

		email := jsondta["email"]

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
			return err, false
		}

		if err != nil {
			fmt.Println("Error generating otp :", err)
			return err, false
		}

		fmt.Println(otpstring)
	}
	return nil, false
}

func Otp_Verify(email, otp string) error {

	fmt.Println(otp)
	fmt.Println("this issssssssssssssssssssssssssssss")

	otpstring, err := otpadapter.GetOTP(email)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("this is the retrieved otp")
	fmt.Println(otpstring)

	if otp != otpstring {
		fmt.Println("otp doesnt match....")
		return fmt.Errorf("otp doesnt match")
	}

	return nil
}
