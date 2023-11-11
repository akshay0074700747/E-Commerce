package usecases

import (
	"context"
	"ecommerce/internal/interfaces/repositories"
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
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

	return coupon.CouponRepo.AddCoupon(req)

}

func (coupon *CouponUsecase) GetAllCouponsByEmail(ctx context.Context, email string) ([]responce.CouponData, error) {

	return coupon.CouponRepo.GetAllCouponsByEmail(email)

}

func (coupon *CouponUsecase) GetAllCoupons(ctx context.Context) ([]responce.CouponData, error) {

	return coupon.CouponRepo.GetAllCoupons()

}

func (coupon *CouponUsecase) GetCouponByID(ctx context.Context, id uint) (responce.CouponData, error) {

	return coupon.CouponRepo.GetCouponByID(id)

}

func (coupon *CouponUsecase) RemoveCouponFromUser(ctx context.Context, id uint, email string) error {

	return coupon.CouponRepo.RemoveCouponFromUser(id,email)

}

func (coupon *CouponUsecase) UpdateCoupon(ctx context.Context, req helperstructs.CouponReq) error {

	return coupon.CouponRepo.UpdateCoupon(req)

}

func (coupon *CouponUsecase) DeleteCoupon(ctx context.Context, id uint) error {

	return coupon.CouponRepo.DeleteCoupon(id)

}
