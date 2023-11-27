package responce

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type JSONB map[string]interface{}

func (j *JSONB) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}

	err := json.Unmarshal(source, j)
	if err != nil {
		return err
	}

	return nil
}

func (j JSONB) Value() (driver.Value, error) {
	jByte, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}

	return jByte, nil
}

type ProuctData struct {
	ID               uint     `json:"id"`
	Category         string   `json:"category"`
	SubCategory      string   `json:"subcategory"`
	Brand            string   `json:"brand"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Price            int      `json:"original_price"`
	Rating           float32  `json:"rating"`
	DiscountedPrice  int      `json:"discounted_price"`
	Stock            int      `json:"stock"`
	Specifications   JSONB    `json:"specifications"`
	Images           []string `json:"images"`
	IsActive         bool     `json:"is_active"`
	RelativeProducts *string  `json:"relative_products"`
	UpdatedBy        string   `json:"updated_by"`
	UpdatedAt        time.Time
	CartandWishList
}

type CartandWishList struct {
	AddedToCart     bool `json:"added_to_cart"`
	AddedToWishList bool `json:"added_to_wishlist"`
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

type CouponData struct {
	ID                   uint   `json:"id"`
	Code                 int    `json:"code"`
	OFF                  int    `json:"off"`
	GiveOnPurchaseAbove  int    `json:"give_onpurchase_above"`
	ApplyOnPurchaseAbove int    `json:"apply_onpurchase_above"`
	IsWelcome            bool   `json:"is_welcome"`
	Description          string `json:"description"`
}
