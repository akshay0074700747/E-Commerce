package entities

type Products struct {
	ID uint `gorm:"primaryKey;unique;not null"`
	Category    Categories `gorm:"foreignKey:categories_id"`
	Brand Brands `gorm:"foreignKey:brands_id"`
	Name string
	Description string
	RelatedProducts []uint
}

func (product *Products) Migrate_me() {
}

type Product_Details struct {
	ID uint `gorm:"primaryKey;unique;not null"`
	Product Products `gorm:"foreignKey:products_id"`
	Price int
	Stock int
	IsAcive bool `gorm:"default:true"`
}

func (details *Product_Details) Migrate_me() {
}