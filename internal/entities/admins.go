package entities

import "time"

type Admins struct {
	ID          uint      `gorm:"primaryKey;unique;not null"`
	Email       string    `json:"email" binding:"required,email" gorm:"unique;not null"`
	Password    string    `json:"password" gorm:"not null"`
	Name        string    `json:"name" binding:"required"`
	Isblocked   bool      `json:"isblocked" gorm:"default:false"`
	UnblockTime time.Time `json:"unblocktime"`
	CreatedAt   time.Time
}

func (admin *Admins) Migrate_me() {
}
