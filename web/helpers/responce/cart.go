package responce

import "time"

type CartData struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CartItemData struct {
	ID        uint      `json:"id"`
	CartID    uint      `json:"cart_id"`
	ProductId uint      `json:"product_id"`
	Quantity  int       `json:"quantity"`
	AddedAt   time.Time `json:"added_at"`
}
