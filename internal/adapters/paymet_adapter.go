package adapters

import (
	"ecommerce/internal/interfaces/repositories"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"

	"gorm.io/gorm"
)

type PaymentAdapter struct {
	DB *gorm.DB
}

func NewPaymentAdapter(db *gorm.DB) repositories.PaymentRepo {
	return &PaymentAdapter{DB: db}
}

func (payment *PaymentAdapter) Makepayment(req helperstructs.PaymentReq) error {

	query := `INSERT INTO payment_details (order_id,email,payment_status,order_total,payment_ref) VALUES($1,$2,$3,$4,$5)`

	return payment.DB.Exec(query, req.OrderID, req.Email, req.PaymentStatus, req.TotalPrice, req.PaymentRef).Error

}

func (payment *PaymentAdapter) GetPaymentDetails(orderid uint) (responce.PaymentData, error) {

	var paymentdta responce.PaymentData

	query := `SELECT * FROM payment_details WHERE order_id = $1`

	return paymentdta, payment.DB.Raw(query, orderid).Scan(&paymentdta).Error

}
