package helperstructs

import "time"

type OrderReq struct {
	ID                 uint      `json:"id"`
	Email              string    `json:"email"`
	AddrID             uint      `json:"addr_id"`
	OrderDate          time.Time `json:"order_date"`
	ExpectedDeliveryBy time.Time `json:"expected_delivery"`
	ShipmentDate       time.Time `json:"shipment_date"`
	RecievedPayment    bool
	IsCancelled        bool `gorm:"default:false"`
	COD                bool `json:"cod"`
	Price              int  `json:"price"`
	ReturnStatus       bool `gorm:"default:false"`
}

type OrderItemReq struct {
	ID           uint `json:"id"`
	OrderId      uint `json:"order_id"`
	ProductId    uint `json:"product_id"`
	Quantity     int  `json:"quantity"`
	ReturnStatus bool `json:"return_status"`
}

type OrderCancel struct {
	OrderID uint `json:"order_id"`
}

type OrderStatus struct {
	OrderID    uint `json:"order_id"`
	StatusCode int  `json:"status_code"`
}

type PaymentReq struct {
	ID            uint      `json:"id"`
	Email         string    `json:"email"`
	OrderID       uint      `json:"order_id"`
	PaymentStatus string    `json:"payment_status"`
	PaymentRef    string    `json:"payment_ref"`
	TotalPrice    int       `json:"total_price"`
	UpdatedAt     time.Time `json:"updated_at"`
}
