package entities

import "time"

type Review struct {
	ID          uint `gorm:"primaryKey;unique;not null"`
	Product     uint
	ReviewedBy  string `gorm:"not null"`
	Description string
	Rating      float32
	CreatedAt   time.Time
}

func (review *Review) Migrate_me() {
}
