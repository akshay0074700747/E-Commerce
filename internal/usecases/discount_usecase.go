package usecases

import (
	"context"
	"ecommerce/internal/interfaces/repositories"
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"strconv"
)

type DiscountUsecase struct {
	DiscountRepo repositories.DiscountRepo
}

func NewDiscountUsecase(repo repositories.DiscountRepo) usecasesinterface.DiscountUsecaseInterface {

	return &DiscountUsecase{DiscountRepo: repo}

}

func (discount *DiscountUsecase) AddDiscount(ctx context.Context, discountreq helperstructs.DiscountReq) (responce.DiscountData, error) {

	return discount.DiscountRepo.AddDiscount(discountreq)

}

func (discount *DiscountUsecase) UpdateDiscount(ctx context.Context, discountreq helperstructs.DiscountReq) (responce.DiscountData, error) {

	return discount.DiscountRepo.UpdateDiscount(discountreq)

}

func (discount *DiscountUsecase) DeleteDiscount(ctx context.Context, id string) error {

	iduint, err := strconv.ParseUint(id, 10, 0)

	if err != nil {
		return err
	}

	return discount.DiscountRepo.DeleteDiscount(uint(iduint))

}

func (discount *DiscountUsecase) GetAllDiscounts(ctx context.Context) ([]responce.DiscountData, error) {

	return discount.DiscountRepo.GetAllDiscounts()

}

func (discount *DiscountUsecase) GetByID(ctx context.Context, category_id string) (responce.DiscountData, error) {

	iduint, err := strconv.ParseUint(category_id, 10, 0)

	if err != nil {
		return responce.DiscountData{}, err
	}

	return discount.DiscountRepo.GetByID(uint(iduint))

}
