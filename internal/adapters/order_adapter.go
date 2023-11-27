package adapters

import (
	"ecommerce/internal/interfaces/repositories"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type OrderAdapter struct {
	DB *gorm.DB
}

func NewOrderAdapter(db *gorm.DB) repositories.OrderRepo {

	return &OrderAdapter{DB: db}

}

func (order *OrderAdapter) AddOrder(ordereq helperstructs.OrderReq) (responce.OrderData, error) {

	fmt.Println(ordereq)

	var orderdata responce.OrderData

	query := `INSERT INTO orders (email,addr_id,expected_delivery_by,cod,price,shipment_date,using_wallet) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id,email,addr_id,expected_delivery_by,cod,price,shipment_date`

	return orderdata, order.DB.Raw(query, ordereq.Email, ordereq.AddrID, ordereq.ExpectedDeliveryBy, ordereq.COD, ordereq.Price, time.Now().Add(24*time.Hour), ordereq.UsingWallet).Scan(&orderdata).Error

}

func (order *OrderAdapter) AddOrderItem(ordereq helperstructs.OrderItemReq) error {

	query := `INSERT INTO order_items (order_id,product_id,quantity) VALUES($1,$2,$3)`

	return order.DB.Exec(query, ordereq.OrderId, ordereq.ProductId, ordereq.Quantity).Error

}

func (order *OrderAdapter) CancelOrder(orderid uint) error {

	query := `UPDATE orders SET is_cancelled = true WHERE id = $1`

	return order.DB.Exec(query, orderid).Error

}

func (order *OrderAdapter) ReturnOrder(orderid uint) error {

	query := `UPDATE orders SET return_status = true WHERE id = $1 AND expected_delivery_by < NOW()`

	result := order.DB.Exec(query, orderid)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("the order cannot be returned , order not recieved yet")
	}

	return nil

}

func (order *OrderAdapter) TruncateOrderItems(order_id uint) ([]responce.OrderItemData, error) {

	var items []responce.OrderItemData

	query := `DELETE FROM order_items WHERE order_id = $1 RETURNING product_id,quantity`

	return items, order.DB.Raw(query, order_id).Scan(&items).Error

}

func (order *OrderAdapter) GetPriceByID(orderid uint) (int, error) {

	var price int

	query := `SELECT price FROM orders WHERE id = $1`

	return price, order.DB.Raw(query, orderid).Scan(&price).Error

}

func (order *OrderAdapter) GetCodById(orderid uint) (bool, error) {

	var cod bool

	query := `SELECT cod FROM orders WHERE id = $1`

	return cod, order.DB.Raw(query, orderid).Scan(&cod).Error

}

func (order *OrderAdapter) GetAllOrders(offset, countt int) ([]responce.OrderData, error) {

	var orderdata []responce.OrderData

	query := `SELECT * FROM orders OFFSET $1 LIMIT $2`

	return orderdata, order.DB.Raw(query, offset, countt).Scan(&orderdata).Error

}

func (order *OrderAdapter) GetAllOrdersByEmail(email string) ([]responce.OrderData, error) {

	var orderdata []responce.OrderData

	query := `SELECT * FROM orders WHERE email = $1`

	return orderdata, order.DB.Raw(query, email).Scan(&orderdata).Error

}

func (order *OrderAdapter) GetEmailByID(order_id uint) (string, error) {

	var email string

	query := `SELECT email FROM orders WHERE id = $1`

	return email, order.DB.Raw(query, order_id).Scan(&email).Error

}

func (order *OrderAdapter) ChangeStatus(order_id uint, status string) error {

	query := `UPDATE orders SET status = $1 WHERE id = $2`

	return order.DB.Exec(query, status, order_id).Error

}

func (order *OrderAdapter) GetAllProductsByOrderID(orderid uint) ([]uint, error) {

	var ids []uint

	query := `SELECT product_id FROM order_items WHERE order_id = $1`

	return ids, order.DB.Raw(query, orderid).Scan(&ids).Error

}

func (order *OrderAdapter) GetOrderByID(orderid uint) (responce.OrderData, error) {

	var data responce.OrderData

	query := `SELECT * FROM orders where id = $1`

	return data, order.DB.Raw(query, orderid).Scan(&data).Error

}

func (order *OrderAdapter) ToggleReceivedPayment(order_id uint) error {

	query := `UPDATE orders SET recieved_payment = true WHERE id = $1`

	return order.DB.Exec(query, order_id).Error

}

func (order *OrderAdapter) CheckRecievedPayment(order_id uint) (bool, error) {

	var rec bool

	query := `SELECT recieved_payment FROM orders WHERE id = $1 AND cod = false`

	return rec, order.DB.Raw(query, order_id).Scan(&rec).Error

}

func (order *OrderAdapter) UpdatePriceByID(orderid uint, price int) error {

	qeuery := `UPDATE orders SET price = $1 WHERE id = $2`

	return order.DB.Exec(qeuery, price, orderid).Error

}
