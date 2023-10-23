package repositories

import (
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type BrandRepo interface {
	GetallBrand() ([]responce.BrandData, error)
	CreateBrand(brandreq helperstructs.BrandReq) (responce.BrandData, error)
	UpdateBrand(brandreq helperstructs.BrandReq) (responce.BrandData, error)
	DeleteBrand(brand_id uint) error
}
