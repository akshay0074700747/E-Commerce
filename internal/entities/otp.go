package entities

import "time"

type OTP struct {
	Value     string
	Email     string
	ExpiresAt time.Duration
}
