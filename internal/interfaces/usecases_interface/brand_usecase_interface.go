package usecasesinterface

import (
	"context"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type BrandUsecaseInterface interface {
	GetallBrand(ctx context.Context) ([]responce.BrandData, error)
	CreateBrand(ctx context.Context, brandreq helperstructs.BrandReq) (responce.BrandData, error)
	UpdateBrand(ctx context.Context, brandreq helperstructs.BrandReq) (responce.BrandData, error)
	DeleteBrand(ctx context.Context, brand_id uint) error
}
