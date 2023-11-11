package entities

import "time"

type Review struct {
	ID          uint `gorm:"primaryKey;unique;not null"`
	Product     uint
	ReviewedBy  string `gorm:"not null"`
	Description string
	Rating      float32 `gorm:"type:float;gorm:ignore_null_properties"`
	CreatedAt   time.Time `gorm:"default:NOW()"`
}

func (review *Review) Migrate_me() {
}
