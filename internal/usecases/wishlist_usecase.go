package usecases

import (
	"context"
	"ecommerce/internal/interfaces/repositories"
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type WishListUseCase struct {
	WishListRepo repositories.WishListRepo
}

func NewWishListUseCase(repo repositories.WishListRepo) usecasesinterface.WishListUseCaseInterface {

	return &WishListUseCase{WishListRepo: repo}

}

func (wishlist *WishListUseCase) CreateWishList(ctx context.Context, email string) error {

	return wishlist.WishListRepo.CreateWishList(email)

}

func (wishlist *WishListUseCase) AddToWishList(ctx context.Context, WishListreq helperstructs.WishListItemsReq) (responce.ProuctData, error) {

	wishlistid, err := wishlist.WishListRepo.GetWishListID(WishListreq.Email)

	if err != nil {
		return responce.ProuctData{}, err
	}

	WishListreq.WishListID = wishlistid

	if err = wishlist.WishListRepo.AddToWishList(WishListreq); err != nil {
		return responce.ProuctData{}, err
	}

	return wishlist.WishListRepo.GetProductByID(WishListreq.ProductId)

}

func (wishlist *WishListUseCase) GetWishListitems(ctx context.Context, email string) ([]responce.ProuctData, error) {

	var productdata []responce.ProuctData

	Wishlistid, err := wishlist.WishListRepo.GetWishListID(email)

	if err != nil {
		return []responce.ProuctData{}, err
	}

	Wishlistitems, err := wishlist.WishListRepo.GetWishListItems(Wishlistid)

	if err != nil {
		return []responce.ProuctData{}, err
	}

	for _, item := range Wishlistitems {

		product, err := wishlist.WishListRepo.GetProductByID(item.ProductId)

		if err != nil {
			return []responce.ProuctData{}, err
		}

		productdata = append(productdata, product)

	}

	return productdata, nil

}
