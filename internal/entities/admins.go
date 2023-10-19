package entities

import "time"

type Admins struct {
	ID uint `gorm:"primaryKey;unique;not null"`
	Email     string `json:"email" binding:"required,email" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	Name      string `json:"name" binding:"required"`
	CreatedAt time.Time
}

func (admin *Admins) Migrate_me() {
}