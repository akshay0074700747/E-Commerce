package entities

type Brands struct {
	ID   uint   `gorm:"primaryKey;unique;not null"`
	Name string `gorm:"not null"`
}

func (brand *Brands) Migrate_me() {
}
