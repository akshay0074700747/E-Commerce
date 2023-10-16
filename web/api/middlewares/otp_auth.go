package middlewares

import (
	"context"
	"ecommerce/internal/adapters"
	"ecommerce/internal/usecases"
	"ecommerce/web/config"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

var otpadapter *adapters.OtpAdapter

func Otp_Gen(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		envs,err := config.LoadEnv("EMAIL","PASSWORD")

		if err != nil {
			panic("Cannot load env files...")
		}

		smtp_username := envs["EMAIL"]
		smtp_password := envs["PASSWORD"]

		smtp_server := "smtp.gmail.com"
		smtp_port := "587"

		var jsondta map[string]interface{}

		if err := c.Bind(&jsondta); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid json data..."})
		}

		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

		ctx := context.Background()

		pong, err := client.Ping(ctx).Result()

		if err != nil {
			fmt.Println("Error connecting to Redis:", err)
		}

		fmt.Println(pong)

		email := jsondta["email"].(string)

		smtpconfig := adapters.SMTPConfig{
			SMTPServer:   smtp_server,
			SMTPPort:     smtp_port,
			SMTPUsername: smtp_username,
			SMTPPassword: smtp_password,
			Receiver: email,
		}

		otpadapter = adapters.NewOtpAdapter(client, smtpconfig)

		otp := usecases.NewOtpUseCase(otpadapter)

		otpstring, err := otp.GenerateOtp(email)

		if err := otp.OTPRepository.SendOTP(otpstring); err != nil {
			fmt.Println(err.Error())
		}

		if err != nil {
			fmt.Println("Error generating otp :", err)
		}

		fmt.Println(otpstring)

		return next(c)
	}
}

func Otp_Verify(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		var jsondta map[string]interface{}

		if err := c.Bind(&jsondta); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid json data..."})
		}

		email, otp := jsondta["email"], jsondta["otp"]

		otpstring,err := otpadapter.GetOTP(email.(string))

		if err != nil {
			fmt.Println(err.Error())
		}

		if otp.(string) != otpstring {
			fmt.Println("otp doesnt match....")
		}

		return next(c)

	}
}
