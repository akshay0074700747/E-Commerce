package usecasesinterface

import (
	"context"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type CouponUsecaseInterface interface {
	AddCoupon(ctx context.Context, req helperstructs.CouponReq) error
	GetAllCouponsByEmail(ctx context.Context, email string) ([]responce.CouponData, error)
	GetAllCoupons(ctx context.Context) ([]responce.CouponData, error)
	GetCouponByCode(ctx context.Context, code int) (responce.CouponData, error)
	RemoveCouponFromUser(ctx context.Context, id uint, email string) error
	UpdateCoupon(ctx context.Context, req helperstructs.CouponReq) error
	DeleteCoupon(ctx context.Context, id uint) error
	ListofCouponsAvailableForThisOrder(price int) ([]uint, error)
	CreditUserWithCoupon(email string, id uint) error
}
