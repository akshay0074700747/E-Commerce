package usecases

import (
	"context"
	"ecommerce/internal/interfaces/repositories"
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	"ecommerce/web/helpers"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"fmt"
)

type OrderUsecase struct {
	OrderRepo    repositories.OrderRepo
	CartRepo     repositories.CartRepo
	ProductRepo  repositories.ProductsRepo
	UserRepo     repositories.UserRepo
	DiscountRepo repositories.DiscountRepo
}

func NewOrderUsecase(orderrepo repositories.OrderRepo, cartrepo repositories.CartRepo, productrepo repositories.ProductsRepo, userrepo repositories.UserRepo, discountrepo repositories.DiscountRepo) usecasesinterface.OrderUsecaseInterface {

	return &OrderUsecase{
		OrderRepo:    orderrepo,
		CartRepo:     cartrepo,
		ProductRepo:  productrepo,
		UserRepo:     userrepo,
		DiscountRepo: discountrepo,
	}

}

func (order *OrderUsecase) AddOrder(ctx context.Context, orderreq helperstructs.OrderReq) (responce.OrderData, error) {

	var orderitmreq helperstructs.OrderItemReq

	cart_id, err := order.CartRepo.GetCartID(orderreq.Email)

	if err != nil {
		return responce.OrderData{}, err
	}

	items, err := order.CartRepo.GetCartitems(cart_id)

	if err != nil {
		return responce.OrderData{}, err
	}

	if orderreq.COD {
		orderreq.RecievedPayment = false
	} else {
		orderreq.RecievedPayment = true
	}

	for i := range items {

		price, err := order.ProductRepo.GetPriceByID(items[i].ProductId)

		if err != nil {
			return responce.OrderData{}, err
		}

		discdta, err := order.DiscountRepo.GetByProductID(items[i].ProductId)

		if err != nil {
			return responce.OrderData{}, err
		}

		fmt.Println(discdta.Discount)

		if discdta.Discount != 0 {
			price -= ((price * int(discdta.Discount)) / 100)
		}

		price *= items[i].Quantity

		if err := order.ProductRepo.DecreaseStock(items[i].ProductId, items[i].Quantity); err != nil {
			return responce.OrderData{}, err
		}

		orderreq.Price += price
	}

	orderreq.ExpectedDeliveryBy = helpers.RandomExpiry()

	orderdta, err := order.OrderRepo.AddOrder(orderreq)

	if err != nil {
		return orderdta, err
	}

	for _, item := range items {

		orderitmreq.OrderId = orderdta.ID
		orderitmreq.ProductId = item.ProductId
		orderitmreq.Quantity = item.Quantity

		if err := order.OrderRepo.AddOrderItem(orderitmreq); err != nil {
			return orderdta, err
		}

	}

	if err = order.CartRepo.TruncateCart(cart_id); err != nil {
		return orderdta, err
	}

	return orderdta, nil

}

func (order *OrderUsecase) CancelOrder(ctx context.Context, orderid uint, email string) error {

	if err := order.OrderRepo.CancelOrder(orderid); err != nil {
		return err
	}

	itemdata, err := order.OrderRepo.TruncateOrderItems(orderid)

	if err != nil {
		return err
	}

	cartid, err := order.CartRepo.GetCartID(email)

	var cartreq helperstructs.CartItemReq

	cartreq.CartID = cartid

	if err != nil {
		return err
	}

	for _, item := range itemdata {

		if err := order.ProductRepo.IncreaseStock(item.ProductID, item.Quantity); err != nil {
			return err
		}

		cartreq.ProductId = item.ProductID

		cartreq.Quantity = item.Quantity

		if err := order.CartRepo.AddToCart(cartreq); err != nil {
			return err
		}

	}

	cod, err := order.OrderRepo.GetCodById(orderid)

	if err != nil {
		return err
	}

	if !cod {
		price, err := order.OrderRepo.GetPriceByID(orderid)

		if err != nil {
			return err
		}

		if err := order.UserRepo.IncrementWallet(email, price); err != nil {
			return err
		}

	}

	return nil

}

func (order *OrderUsecase) ReturnOrder(ctx context.Context, orderid uint, email string) error {

	if err := order.OrderRepo.ReturnOrder(orderid); err != nil {
		return err
	}

	items, err := order.OrderRepo.TruncateOrderItems(orderid)

	if err != nil {
		return err
	}

	for _, item := range items {

		if err := order.ProductRepo.IncreaseStock(item.ProductID, item.Quantity); err != nil {
			return err
		}

	}

	price, err := order.OrderRepo.GetPriceByID(orderid)

	if err != nil {
		return err
	}

	if err := order.UserRepo.IncrementWallet(email, price); err != nil {
		return err
	}

	return nil

}

func (order *OrderUsecase) GetAllOrders(ctx context.Context) ([]responce.OrderData, error) {

	return order.OrderRepo.GetAllOrders()

}

func (order *OrderUsecase) GetAllOrdersByEmail(ctx context.Context, email string) ([]responce.OrderData, error) {

	orderdta, err := order.OrderRepo.GetAllOrdersByEmail(email)

	if err != nil {
		return orderdta, err
	}

	for i := range orderdta {

		orderdta[i].Products, err = order.OrderRepo.GetAllProductsByOrderID(orderdta[i].ID)

		if err != nil {
			return orderdta, err
		}

	}

	return orderdta, nil

}

func (order *OrderUsecase) GetEmailByID(ctx context.Context, order_id uint) (string, error) {

	return order.OrderRepo.GetEmailByID(order_id)

}

func (order *OrderUsecase) ChangeStatus(ctx context.Context, order_id uint, status string) error {

	return order.OrderRepo.ChangeStatus(order_id, status)

}

func (order *OrderUsecase) GetOrderByID(orderid uint) (responce.OrderData, error) {

	return order.OrderRepo.GetOrderByID(orderid)

}

func (order *OrderUsecase) ToggleReceivedPayment(order_id uint) error {

	return order.OrderRepo.ToggleReceivedPayment(order_id)

}
