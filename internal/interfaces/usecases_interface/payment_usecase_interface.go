package usecasesinterface

import (
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type PaymentUsecaseInterface interface{
	Makepayment(req helperstructs.PaymentReq) error
	GetPaymentDetails(orderid uint) (responce.PaymentData,error)
}