package entities

import "time"

type PaymentDetails struct {
	ID            uint `gorm:"primaryKey"`
	Email         string
	OrderID       uint
	OrderTotal    float64
	PaymentStatus string
	PaymentRef    string
	UpdatedAt     time.Time
}

func (payment *PaymentDetails) Migrate_me() {
}
