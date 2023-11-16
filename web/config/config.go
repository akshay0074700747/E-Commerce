package config

import (
	"os"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
)

type Config struct {
	DATABASE_ADDR  string `validate:"required"`
	SECRET         string `validate:"required"`
	EMAIL          string `validate:"required"`
	PASSWORD       string `validate:"required"`
	RAZORPAYID     string `validate:"required"`
	RAZORPAYSECRET string `validate:"required"`
}

func LoadConfig() (Config, error) {

	var config Config

	if err := godotenv.Load(".env"); err != nil {
		return config, err
	}

	// Read environment variables
	config.DATABASE_ADDR = os.Getenv("DATABASE_ADDR")
	config.SECRET = os.Getenv("SECRET")
	config.EMAIL = os.Getenv("EMAIL")
	config.PASSWORD = os.Getenv("PASSWORD")
	config.RAZORPAYID = os.Getenv("RAZORPAYID")
	config.RAZORPAYSECRET = os.Getenv("RAZORPAYSECRET")

	// Perform validation
	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return config, err
	}

	return config, nil
}
