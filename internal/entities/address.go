package entities

type Address struct {
	ID            uint   `gorm:"primaryKey;unique;not null"`
	Email         string `gorm:"not null"`
	HouseName     string `gorm:"not null"`
	StreetAddress string `gorm:"not null"`
	City          string `gorm:"not null"`
	District      string `gorm:"not null"`
	PO            string `gorm:"not null"`
	State         string `gorm:"not null"`
}

func (address *Address) Migrate_me() {
}
