package usecasesinterface

import (
	"context"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type OrderUsecaseInterface interface {
	AddOrder(context.Context, helperstructs.OrderReq) (responce.OrderData, error)
	CancelOrder(ctx context.Context, orderid uint, email string) error
	ReturnOrder(ctx context.Context, orderid uint, email string) error
	GetAllOrders(ctx context.Context,count,page string) ([]responce.OrderData, error)
	GetAllOrdersByEmail(ctx context.Context, email string) ([]responce.OrderData, error)
	GetEmailByID(ctx context.Context, order_id uint) (string, error)
	ChangeStatus(ctx context.Context, order_id uint, status string) error
	GetOrderByID(orderid uint) (responce.OrderData, error)
	ToggleReceivedPayment(order_id uint) error
}
