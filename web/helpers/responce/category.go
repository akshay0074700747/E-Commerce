package responce

import "time"

type CategoryData struct {
	ID uint `json:"id"`
	Category string `json:"category"`
	SubCategory string `json:"subcategory"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
	UpdatedBy string `json:"updatedby"`
}