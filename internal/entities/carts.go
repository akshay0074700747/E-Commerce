package entities

import "time"

type Carts struct {
	ID        uint   `gorm:"primaryKey;unique;not null"`
	Email     string `gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (cart *Carts) Migrate_me() {
}

type CartItems struct {
	ID        uint `gorm:"primaryKey;unique;not null"`
	CartID    uint `gorm:"not null"`
	ProductId uint `gorm:"not null"`
	Quantity  int  `gorm:"default:1"`
	AddedAt   time.Time
}

func (cartitem *CartItems) Migrate_me() {
}
