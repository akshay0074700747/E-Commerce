package usecases

import (
	"context"
	"ecommerce/internal/interfaces/repositories"
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	"ecommerce/web/helpers"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type CouponUsecase struct {
	CouponRepo repositories.Coupons
}

func NewCouponUsecase(repo repositories.Coupons) usecasesinterface.CouponUsecaseInterface {
	return &CouponUsecase{CouponRepo: repo}
}

func (coupon *CouponUsecase) AddCoupon(ctx context.Context, req helperstructs.CouponReq) error {

	req.Code = helpers.SelectRandomintBetweenRange(100000,999999)

	return coupon.CouponRepo.AddCoupon(req)

}

func (coupon *CouponUsecase) GetAllCouponsByEmail(ctx context.Context, email string) ([]responce.CouponData, error) {

	return coupon.CouponRepo.GetAllCouponsByEmail(email)

}

func (coupon *CouponUsecase) GetAllCoupons(ctx context.Context) ([]responce.CouponData, error) {

	return coupon.CouponRepo.GetAllCoupons()

}

func (coupon *CouponUsecase) GetCouponByCode(ctx context.Context, code int) (responce.CouponData, error) {

	return coupon.CouponRepo.GetCouponByCode(code)

}

func (coupon *CouponUsecase) RemoveCouponFromUser(ctx context.Context, id uint, email string) error {

	return coupon.CouponRepo.RemoveCouponFromUser(id, email)

}

func (coupon *CouponUsecase) UpdateCoupon(ctx context.Context, req helperstructs.CouponReq) error {

	return coupon.CouponRepo.UpdateCoupon(req)

}

func (coupon *CouponUsecase) DeleteCoupon(ctx context.Context, id uint) error {

	return coupon.CouponRepo.DeleteCoupon(id)

}

func (coupon *CouponUsecase) ListofCouponsAvailableForThisOrder(price int) ([]uint, error) {

	return coupon.CouponRepo.ListofCouponsAvailableForThisOrder(price)

}

func (coupon *CouponUsecase) CreditUserWithCoupon(email string, id uint) error {

	return coupon.CouponRepo.CreditUserWithCoupon(email, id)

}
