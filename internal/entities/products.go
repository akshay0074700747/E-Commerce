package entities

import "time"

type RelatedProductSlice []int

type Products struct {
	ID               uint `gorm:"primaryKey;unique;not null"`
	Category         uint
	Brand            uint
	Name             string
	Description      string
	Price            int
	Stock            int
	Specifications   map[string]interface{} `gorm:"type:jsonb"`
	IsActive         bool                   `gorm:"default:true"`
	RelativeProducts []uint
	UpdatedAt        time.Time
	UpdatedBy        string `gorm:"not null"`
}

type Images struct {
	ID        uint `gorm:"primaryKey;unique;not null"`
	ProductID uint
	Image     string
}

func (images *Images) Migrate_me() {
}

func (product *Products) Migrate_me() {
}
