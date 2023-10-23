package entities

import "time"

type Categories struct {
	ID          uint   `gorm:"primaryKey;unique;not null"`
	Category    string `gorm:"not null"`
	SubCategory string `gorm:"not null"`
	UpdatedAt   time.Time
	UpdatedBy   string `gorm:"not null"`
}

func (category *Categories) Migrate_me() {
}
