package usecases

import (
	"context"
	"ecommerce/internal/interfaces/repositories"
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"strconv"
)

type ProductUsecases struct {
	ProductRepo repositories.ProductsRepo
}

func NewProductUsecases(repo repositories.ProductsRepo) usecasesinterface.ProductUsecaseInterface {
	return &ProductUsecases{ProductRepo: repo}
}

func (product *ProductUsecases) AddProduct(ctx context.Context, productreq helperstructs.ProductReq) (responce.ProuctData, error) {

	if productreq.RelatedProducts == nil {

		var err error

		productreq.RelatedProducts, err = product.ProductRepo.FindRelatedProducts(productreq.Category)

		if err != nil {
			return responce.ProuctData{}, err
		}

	}

	return product.ProductRepo.AddProduct(productreq)

}

func (product *ProductUsecases) GetProducts(ctx context.Context, email string) ([]responce.ProuctData, error) {

	productsdata, err := product.ProductRepo.GetProducts()

	if err != nil {
		return productsdata, err
	}

	if email != "" {
		
	}

	for i := range productsdata {

		brand, err := strconv.ParseUint(productsdata[i].Brand, 10, 0)

		if err != nil {
			return productsdata, err
		}

		productsdata[i].Brand, err = product.ProductRepo.GetBrand(uint(brand))

		if err != nil {
			return productsdata, err
		}

		cat, err := product.ProductRepo.GetCategoryID(productsdata[i].Category, productsdata[i].SubCategory)

		if err != nil {
			return productsdata, err
		}

		discount, err := product.ProductRepo.FindDiscountByID(cat)

		if err != nil {
			return productsdata, err
		}
		if discount.Discount != float32(0) {
			productsdata[i].DiscountedPrice = productsdata[i].Price - int((discount.Discount*float32(productsdata[i].Price))/100)
		}

	}

	return productsdata, nil

}

func (product *ProductUsecases) UpdateProduct(ctx context.Context, productreq helperstructs.ProductReq) (responce.ProuctData, error) {

	return product.ProductRepo.UpdateProduct(productreq)

}

func (product *ProductUsecases) DeleteProduct(ctx context.Context, product_id uint) error {

	return product.ProductRepo.DeleteProduct(product_id)

}

func (product *ProductUsecases) UpdateStock(ctx context.Context, id uint, stock int) error {

	return product.ProductRepo.UpdateStock(id, stock)

}

func (product *ProductUsecases) GetProductByID(ctx context.Context, id string) (responce.ProuctData, error) {

	uintid, err := strconv.ParseUint(id, 10, 0)

	if err != nil {
		return responce.ProuctData{}, err
	}

	return product.ProductRepo.GetProductByID(uint(uintid))

}
