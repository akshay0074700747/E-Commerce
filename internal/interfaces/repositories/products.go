package repositories

import (
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type ProductsRepo interface {
	AddProduct(productreq helperstructs.ProductReq) (responce.ProuctData, error)
	GetProducts() ([]responce.ProuctData, error)
	UpdateProduct(productreq helperstructs.ProductReq) (responce.ProuctData, error)
	DeleteProduct(product_id uint) error
	FindRelatedProducts(cat_id uint) ([]uint, error)
	FindDiscountByID(category_id uint) (responce.DiscountData, error)
	GetCategoryID(category, subcategory string) (uint, error)
	GetBrand(id uint) (string, error)
	UpdateStock(id uint, stock int) error
	GetProductByID(id uint) (responce.ProuctData, error)
	GetPriceByID(id uint) (int, error)
	IncreaseStock(id uint, stock int) error
	DecreaseStock(id uint, stock int) error
}
