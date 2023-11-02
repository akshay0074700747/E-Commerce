package usecasesinterface

import (
	"context"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type CartUseCaseInterface interface {
	CreateCart(ctx context.Context, email string) error
	AddToCart(ctx context.Context, cartreq helperstructs.CartItemReq) (responce.ProuctData, error)
	GetCartitems(ctx context.Context, email string) ([]responce.ProuctData, error)
	UpdateQuantity(ctx context.Context, cartreq helperstructs.CartItemReq) error
	DeleteCartItem(ctx context.Context, cartreq helperstructs.CartItemReq) error
}
