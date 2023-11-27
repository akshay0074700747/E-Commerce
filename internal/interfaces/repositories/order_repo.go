package repositories

import (
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type OrderRepo interface {
	AddOrder(helperstructs.OrderReq) (responce.OrderData, error)
	AddOrderItem(helperstructs.OrderItemReq) error
	CancelOrder(orderid uint) error
	ReturnOrder(orderid uint) error
	TruncateOrderItems(order_id uint) ([]responce.OrderItemData, error)
	GetPriceByID(orderid uint) (int, error)
	GetCodById(orderid uint) (bool, error)
	GetAllOrders(offset,countt int) ([]responce.OrderData, error)
	GetAllOrdersByEmail(email string) ([]responce.OrderData, error)
	GetEmailByID(order_id uint) (string, error)
	ChangeStatus(order_id uint, status string) error
	GetAllProductsByOrderID(orderid uint) ([]uint, error)
	GetOrderByID(orderid uint) (responce.OrderData, error)
	ToggleReceivedPayment(order_id uint) error
	CheckRecievedPayment(order_id uint) (bool, error)
	UpdatePriceByID(orderid uint, price int) error
}
