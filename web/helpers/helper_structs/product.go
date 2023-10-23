package helperstructs

type ProductReq struct {
	ID              uint                   `json:"id"`
	Category        uint                   `json:"category_id"`
	Brand           uint                   `json:"brand"`
	Name            string                 `json:"name"`
	Description     string                 `json:"description"`
	Price           int                    `json:"price"`
	Stock           int                    `json:"stock"`
	Specifications  map[string]interface{} `json:"specifications"`
	IsActive        bool                   `json:"is_active"`
	RelatedProducts []uint                 `json:"related_products"`
	UpdatedBy       string                 `json:"updated_by"`
}

type BrandReq struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
