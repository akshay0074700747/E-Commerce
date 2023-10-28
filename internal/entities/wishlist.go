package entities

type WishLists struct {
	ID    uint   `gorm:"primaryKey;unique;not null"`
	Email string `gorm:"unique;not null"`
}

func (wish *WishLists) Migrate_me() {
}

type WishListItems struct {
	ID         uint `gorm:"primaryKey;unique;not null"`
	WishListID uint `gorm:"not null"`
	ProductId  uint `gorm:"not null"`
}

func (wishlist *WishListItems) Migrate_me() {
}
