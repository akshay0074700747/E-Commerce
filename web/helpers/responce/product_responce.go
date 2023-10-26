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
		return errors.New("Type assertion .([]byte) failed.")
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
	ID               uint   `json:"id"`
	Category         string `json:"category"`
	SubCategory      string `json:"subcategory"`
	Brand            string `json:"brand"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Price            int    `json:"original_price"`
	DiscountedPrice  int    `json:"discounted_price"`
	Stock            int    `json:"stock"`
	Specifications   JSONB  `json:"specifications"`
	IsActive         bool   `json:"is_active"`
	RelativeProducts string `json:"relative_products"`
	UpdatedBy        string `json:"updated_by"`
	UpdatedAt        time.Time
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
