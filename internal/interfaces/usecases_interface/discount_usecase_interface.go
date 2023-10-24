package usecasesinterface

import (
	"context"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type DiscountUsecaseInterface interface {
	GetByID(ctx context.Context, category_id string) (responce.DiscountData, error)
	GetAllDiscounts(ctx context.Context) ([]responce.DiscountData, error)
	AddDiscount(ctx context.Context, discountreq helperstructs.DiscountReq) (responce.DiscountData, error)
	UpdateDiscount(ctx context.Context, discountreq helperstructs.DiscountReq) (responce.DiscountData, error)
	DeleteDiscount(ctx context.Context, id string) error
}
