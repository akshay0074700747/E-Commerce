package repositories

import (
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type CartRepo interface {
	CreateCart(email string) error
	AddToCart(cartreq helperstructs.CartItemReq) error
	GetCartitems(id uint) ([]responce.CartItemData, error)
	GetCartID(email string) (uint, error)
	GetProductByID(id uint) (responce.ProuctData, error)
	GetItemByProductID(cart_id, product_id uint) (responce.CartItemData, error)
	UpdateQuantity(cartreq helperstructs.CartItemReq) error
	DeleteCartItem(cartreq helperstructs.CartItemReq) error
	TruncateCart(cart_id uint) ([]responce.CartItemData, error)
}
