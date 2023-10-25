package responce

import "time"

type ProuctData struct {
	ID              uint                   `json:"id"`
	Category        string                 `json:"category"`
	SubCategory     string                 `json:"subcategory"`
	Brand           string                 `json:"brand"`
	Name            string                 `json:"name"`
	Description     string                 `json:"description"`
	Price           int                    `json:"original_price"`
	DiscountedPrice int                    `json:"discounted_price"`
	Stock           int                    `json:"stock"`
	Specifications  map[string]interface{} `json:"specifications"`
	IsActive        bool                   `json:"is_active"`
	RelatedProducts []uint                 `json:"related_products"`
	UpdatedBy       string                 `json:"updated_by"`
	UpdatedAt       time.Time
}

type BrandData struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type DiscountData struct {
	ID        uint      `json:"id"`
	Category  uint      `json:"category_id"`
	Discount  float32   `json:"discount"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
