package responce

import "time"

type ProuctData struct {
	ID              uint                   `json:"id"`
	Category        string                 `json:"category"`
	SubCategory     string                 `json:"subcategory"`
	Brand           string                 `json:"brand"`
	Name            string                 `json:"name"`
	Description     string                 `json:"description"`
	Price           int                    `json:"price"`
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
