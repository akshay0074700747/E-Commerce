package entities

import "time"

type Orders struct {
	ID                    uint      `gorm:"primaryKey;unique;not null"`
	Email                 string    `gorm:"not null"`
	AddrID                uint      `gorm:"not null"`
	OrderDate             time.Time `gorm:"default:NOW()"`
	ExpectedDeliveryBy    time.Time
	RecievedPayment       bool `gorm:"default:false"`
	DecrementedFromWallet int  `gorm:"default:0"`
	IsCancelled           bool `gorm:"default:false"`
	UsingWallet           bool `gorm:"default:false"`
	ShipmentDate          time.Time
	Status                string
	COD                   bool
	Price                 int
	ReturnStatus          bool `gorm:"default:false"`
}

func (order *Orders) Migrate_me() {
}

type OrderItems struct {
	ID        uint `gorm:"primaryKey;unique;not null"`
	OrderId   uint `gorm:"not null"`
	ProductId uint `gorm:"not null"`
	Quantity  int
}

func (order_item *OrderItems) Migrate_me() {
}
