package repositories

import (
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type WishListRepo interface {
	CreateWishList(email string) error
	AddToWishList(wishreq helperstructs.WishListItemsReq) error
	GetWishListItems(id uint) ([]responce.WishListItemsData, error)
	GetWishListID(email string) (uint, error)
	GetProductByID(id uint) (responce.ProuctData, error)
}