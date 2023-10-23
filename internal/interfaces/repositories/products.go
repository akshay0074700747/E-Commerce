package repositories

import (
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type ProductsRepo interface {
	AddProduct(productreq helperstructs.ProductReq) (responce.ProuctData,error)
	GetProducts()([]responce.ProuctData,error)
	UpdateProduct(productreq helperstructs.ProductReq) (responce.ProuctData,error)
	DeleteProduct(product_id uint ) (error)
	FindRelatedProducts(cat_id uint) ([]uint,error)
}