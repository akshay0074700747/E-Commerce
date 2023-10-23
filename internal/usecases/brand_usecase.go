package usecases

import (
	"context"
	"ecommerce/internal/interfaces/repositories"
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type BrandUsecase struct {
	BrandRepo repositories.BrandRepo
}

func NewBrandUsecase(repo repositories.BrandRepo) usecasesinterface.BrandUsecaseInterface {

	return &BrandUsecase{BrandRepo: repo}

}

func (brand *BrandUsecase) CreateBrand(ctx context.Context, brandreq helperstructs.BrandReq) (responce.BrandData, error) {

	return brand.BrandRepo.CreateBrand(brandreq)

}

func (brand *BrandUsecase) GetallBrand(ctx context.Context) ([]responce.BrandData, error) {

	return brand.BrandRepo.GetallBrand()

}

func (brand *BrandUsecase) UpdateBrand(ctx context.Context, brandreq helperstructs.BrandReq) (responce.BrandData, error) {

	return brand.BrandRepo.UpdateBrand(brandreq)

}

func (brand *BrandUsecase) DeleteBrand(ctx context.Context, brand_id uint) error {

	return brand.BrandRepo.DeleteBrand(brand_id)

}
