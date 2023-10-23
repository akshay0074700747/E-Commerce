package entities

import "time"

type Discount struct {
	ID        uint `gorm:"primaryKey;unique;not null"`
	Category  uint
	Discount  float32 `gorm:"not null"`
	StartDate time.Time
	EndDate   time.Time
}

func (discount *Discount) Migrate_me() {
}
