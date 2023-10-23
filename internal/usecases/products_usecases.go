package usecases

import (
	"context"
	"ecommerce/internal/interfaces/repositories"
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
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

func (product *ProductUsecases) GetProducts(ctx context.Context) ([]responce.ProuctData, error) {

	return product.ProductRepo.GetProducts()

}

func (product *ProductUsecases) UpdateProduct(ctx context.Context, productreq helperstructs.ProductReq) (responce.ProuctData, error) {

	return product.ProductRepo.UpdateProduct(productreq)

}

func (product *ProductUsecases) DeleteProduct(ctx context.Context, product_id uint) error {

	return product.ProductRepo.DeleteProduct(product_id)

}
