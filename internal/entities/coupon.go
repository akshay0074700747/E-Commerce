package entities

type Coupon struct {
	ID                   uint `gorm:"primaryKey;unique;not null"`
	OFF                  int
	GiveOnPurchaseAbove  int
	ApplyOnPurchaseAbove int
	Description          string
}

func (coupon *Coupon) Migrate_me() {
}

type CouponItems struct {
	ID     uint `gorm:"primaryKey;unique;not null"`
	Email  string
	Coupon uint
}

func (coupon_items *CouponItems) Migrate_me() {
}
