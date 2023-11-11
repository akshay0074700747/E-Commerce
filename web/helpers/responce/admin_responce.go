package responce

import "time"

type AdminData struct {
	Id        uint   `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Isblocked bool   `json:"isblocked"`
	Reports   int    `json:"reports"`
}

type AdminsideUsersData struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Mobile    string    `json:"mobile"`
	Isblocked bool      `json:"isblocked"`
	Reports   int       `json:"reports"`
	CreatedAt time.Time `json:"createdat"`
}

type AdminDashBoard struct {
	Users                    int    `json:"total_users"`
	Products                 int    `json:"total_products"`
	Sales                    int    `json:"total_sales"`
	CancelledOrders          int    `json:"cancelled_orders"`
	DeliveredOrders          int    `json:"delivered_orders"`
	PurchasedUsers           int    `json:"purchased_users"`
	MostPurchasedCategory    string `json:"bestseller_category"`
	MostPurchasedSubCategory string `json:"bestseller_subcategory"`
	MostPurchasedProduct     string `json:"bestseller_product"`
	ActiveDiscounts          []uint `json:"active_discoounts"`
	TotalBlockedUsers        int    `json:"total_blockedusers"`
}

type OrderedProductWithCount struct {
	ProductID uint `json:"product_id"`
	Count     int  `json:"count"`
}

type AdminSalesReport struct {
	Orders            int `json:"orders"`
	TransactionAmount int `json:"transaction_amount"`
	ProductsSold      int `json:"sold_products"`
	BuyedUsers        int `json:"buyed_users"`
}
