package repositories

import (
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type Coupons interface {
	AddCoupon(req helperstructs.CouponReq) (error)
	GetAllCouponsByEmail(email string) ([]responce.CouponData,error)
	GetAllCoupons() ([]responce.CouponData,error)
	GetCouponByID(id uint) (responce.CouponData,error)
	RemoveCouponFromUser(id uint,email string) (error)
	UpdateCoupon(req helperstructs.CouponReq) (error)
	DeleteCoupon(id uint) (error)
}