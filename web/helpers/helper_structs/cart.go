package helperstructs

type CartReq struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

type CartItemReq struct {
	ID        uint `json:"id"`
	CartID    uint `json:"cart_id"`
	ProductId uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
	Email     string `json:"email"`
}
