package entities

import "time"

type User struct {
	ID          uint      `gorm:"primaryKey;unique;not null"`
	Name        string    `json:"name" binding:"required"`
	Email       string    `json:"email" binding:"required,email" gorm:"unique;not null"`
	Mobile      string    `json:"mobile" binding:"required,eq=10" gorm:"unique; not null"`
	Password    string    `json:"password" gorm:"not null"`
	Wallet      int       `json:"wallet" gorm:"default:0"`
	Isblocked   bool      `json:"isblocked" gorm:"default:false"`
	UnblockTime time.Time `json:"unblocktime"`
	CreatedAt   time.Time
}

func (user *User) Migrate_me() {
}
