package usecases

import (
	"context"
	"ecommerce/internal/interfaces/repositories"
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"fmt"
	"strconv"
)

type ProductUsecases struct {
	ProductRepo  repositories.ProductsRepo
	CartRepo     repositories.CartRepo
	WishListRepo repositories.WishListRepo
}

func NewProductUsecases(repo repositories.ProductsRepo, cartRepo repositories.CartRepo, wishListRepo repositories.WishListRepo) usecasesinterface.ProductUsecaseInterface {
	return &ProductUsecases{ProductRepo: repo, CartRepo: cartRepo, WishListRepo: wishListRepo}
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

			fmt.Println(email)

			cart_id, err := product.CartRepo.GetCartID(email)

			if err != nil {
				return productsdata, err
			}

			item, err := product.CartRepo.GetItemByProductID(cart_id, productsdata[i].ID)

			if err != nil {
				return productsdata, err
			}

			if item.ProductId != 0 {
				productsdata[i].AddedToCart = true
			}

			wish_id, err := product.WishListRepo.GetWishListID(email)

			if err != nil {
				return productsdata, err
			}

			wishitem, err := product.WishListRepo.GetItemByProductID(wish_id, productsdata[i].ID)

			if err != nil {
				return productsdata, err
			}

			if wishitem.ProductId != 0 {
				productsdata[i].AddedToWishList = true
			}

		}

	} else {

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

func (product *ProductUsecases) GetProductByID(ctx context.Context, id string, email string) (responce.ProuctData, error) {

	fmt.Println(id)
	fmt.Println(email)

	uintid, err := strconv.ParseUint(id, 10, 0)

	if err != nil {
		return responce.ProuctData{}, err
	}

	proddata, err := product.ProductRepo.GetProductByID(uint(uintid))

	if err != nil {
		return proddata, err
	}

	if email != "" {

		cart_id, err := product.CartRepo.GetCartID(email)

		if err != nil {
			return proddata, err
		}

		item, err := product.CartRepo.GetItemByProductID(cart_id, proddata.ID)

		if err != nil {
			return proddata, err
		}

		if item.ProductId != 0 {
			proddata.AddedToCart = true
		}

		wish_id, err := product.WishListRepo.GetWishListID(email)

		if err != nil {
			return proddata, err
		}

		wishitem, err := product.WishListRepo.GetItemByProductID(wish_id, proddata.ID)

		if err != nil {
			return proddata, err
		}

		if wishitem.ProductId != 0 {
			proddata.AddedToWishList = true
		}

	}

	return proddata, nil

}
