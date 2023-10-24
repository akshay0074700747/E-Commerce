package repositories

import (
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type DiscountRepo interface {
	GetByID(category_id uint) (responce.DiscountData, error)
	GetAllDiscounts() ([]responce.DiscountData, error)
	AddDiscount(discountreq helperstructs.DiscountReq) (responce.DiscountData, error)
	UpdateDiscount(discountreq helperstructs.DiscountReq) (responce.DiscountData, error)
	DeleteDiscount(id uint) error
}
