package repositories

import (
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type PaymentRepo interface {
	Makepayment(req helperstructs.PaymentReq) error
	GetPaymentDetails(orderid uint) (responce.PaymentData,error)
}
