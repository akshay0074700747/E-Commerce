package usecases

import (
	"ecommerce/internal/interfaces/repositories"
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type PaymentUsecase struct {
	PaymentRepo repositories.PaymentRepo
}

func NewPaymentUsecase(usecase repositories.PaymentRepo) usecasesinterface.PaymentUsecaseInterface {
	return &PaymentUsecase{PaymentRepo: usecase}
}

func (payment *PaymentUsecase) Makepayment(req helperstructs.PaymentReq) error {
	return payment.PaymentRepo.Makepayment(req)
}

func (payment *PaymentUsecase) GetPaymentDetails(orderid uint) (responce.PaymentData, error) {
	return payment.PaymentRepo.GetPaymentDetails(orderid)
}
