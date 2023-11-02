package usecasesinterface

import (
	"context"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type WishListUseCaseInterface interface {
	CreateWishList(ctx context.Context, email string) error
	AddToWishList(ctx context.Context, WishListreq helperstructs.WishListItemsReq) (responce.ProuctData, error)
	GetWishListitems(ctx context.Context, email string) ([]responce.ProuctData, error)
	DeleteWishListItem(ctx context.Context, wishlistreq helperstructs.WishListItemsReq) error
}
