package usecases

import (
	"context"
	"ecommerce/internal/interfaces/repositories"
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"fmt"
)

type CartUseCase struct {
	CartRepo repositories.CartRepo
}

func NewCartUseCase(repo repositories.CartRepo) usecasesinterface.CartUseCaseInterface {

	return &CartUseCase{CartRepo: repo}

}

func (cart *CartUseCase) CreateCart(ctx context.Context, email string) error {

	return cart.CartRepo.CreateCart(email)

}

func (cart *CartUseCase) AddToCart(ctx context.Context, cartreq helperstructs.CartItemReq) (responce.ProuctData, error) {

	cartid, err := cart.CartRepo.GetCartID(cartreq.Email)

	if err != nil {
		return responce.ProuctData{}, err
	}

	cartreq.CartID = cartid

	item, _ := cart.CartRepo.GetItemByProductID(cartreq.CartID, cartreq.ProductId)

	if item.ProductId != 0 {
		return responce.ProuctData{}, fmt.Errorf("the selected item is already present in the cart")
	}

	if err = cart.CartRepo.AddToCart(cartreq); err != nil {
		return responce.ProuctData{}, err
	}

	return cart.CartRepo.GetProductByID(cartreq.ProductId)

}

func (cart *CartUseCase) GetCartitems(ctx context.Context, email string) ([]responce.ProuctData, error) {

	var productdata []responce.ProuctData

	cartid, err := cart.CartRepo.GetCartID(email)

	if err != nil {
		return []responce.ProuctData{}, err
	}

	cartitems, err := cart.CartRepo.GetCartitems(cartid)

	if err != nil {
		return []responce.ProuctData{}, err
	}

	for _, item := range cartitems {

		product, err := cart.CartRepo.GetProductByID(item.ProductId)

		if err != nil {
			return []responce.ProuctData{}, err
		}

		productdata = append(productdata, product)

	}

	return productdata, nil

}

func (cart *CartUseCase) UpdateQuantity(ctx context.Context, cartreq helperstructs.CartItemReq) error {

	cart_id, err := cart.CartRepo.GetCartID(cartreq.Email)

	if err != nil {
		return err
	}

	cartreq.CartID = cart_id

	return cart.CartRepo.UpdateQuantity(cartreq)

}

func (cart *CartUseCase) DeleteCartItem(ctx context.Context, cartreq helperstructs.CartItemReq) error {

	cart_id, err := cart.CartRepo.GetCartID(cartreq.Email)

	if err != nil {
		return err
	}

	cartreq.CartID = cart_id

	return cart.CartRepo.DeleteCartItem(cartreq)

}
