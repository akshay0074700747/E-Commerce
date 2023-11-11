package responce

import "time"

type OrderData struct {
	ID                 uint   `json:"id"`
	Email              string `json:"email"`
	AddrID             uint
	OrderDate          time.Time `json:"order_date"`
	ExpectedDeliveryBy time.Time `json:"expected_delivery_by"`
	ShipmentDate       time.Time `json:"shipment_date"`
	IsShipped          bool      `json:"is_shipped"`
	IsCancelled        bool      `json:"is_cancelled"`
	ReturnStatus       bool      `json:"is_returned"`
	COD                bool      `json:"cod"`
	Price              int       `json:"price"`
	Status             string    `json:"status"`
	Products           []uint    `json:"products"`
}

type OrderItemData struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type PaymentData struct {
	ID            uint      `json:"id"`
	OrderID       uint      `json:"order_id"`
	PaymentStatus string    `json:"payment_status"`
	UpdatedAt     time.Time `json:"updated_at"`
}
