package adapters

import (
	"ecommerce/internal/interfaces/repositories"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"

	"gorm.io/gorm"
)

type CartAdapter struct {
	DB *gorm.DB
}

func NewCartAdapter(db *gorm.DB) repositories.CartRepo {

	return &CartAdapter{DB: db}

}

func (cart *CartAdapter) CreateCart(email string) error {

	query := `INSERT INTO carts (email,created_at) VALUES($1,NOW())`

	return cart.DB.Exec(query, email).Error

}

func (cart *CartAdapter) AddToCart(cartreq helperstructs.CartItemReq) error {

	query := `INSERT INTO cart_items (cart_id,product_id,quantity,added_at) VALUES($1,$2,$3,NOW())`

	return cart.DB.Exec(query, cartreq.CartID, cartreq.ProductId, cartreq.Quantity).Error

}

func (cart *CartAdapter) GetCartitems(id uint) ([]responce.CartItemData, error) {

	var itemdata []responce.CartItemData

	query := `SELECT * FROM cart_items WHERE cart_id = $1`

	return itemdata, cart.DB.Raw(query, id).Scan(&itemdata).Error

}

func (cart *CartAdapter) GetCartID(email string) (uint, error) {

	var id uint

	query := `SELECT id FROM carts WHERE email = $1`

	return id, cart.DB.Raw(query, email).Scan(&id).Error

}

func (cart *CartAdapter) GetProductByID(id uint) (responce.ProuctData, error) {

	var products responce.ProuctData

	query := `SELECT * FROM products WHERE id = $1`

	return products, cart.DB.Raw(query, id).Scan(&products).Error

}

func (cart *CartAdapter) GetItemByProductID(cart_id, product_id uint) (responce.CartItemData, error) {

	var item responce.CartItemData

	query := `SELECT * FROM cart_items WHERE cart_id = $1 AND product_id = $2`

	return item, cart.DB.Raw(query, cart_id, product_id).Scan(&item).Error

}

func (cart *CartAdapter) UpdateQuantity(cartreq helperstructs.CartItemReq) error {

	query := `UPDATE cart_items SET quantity = $1 WHERE cart_id = $2 AND product_id = $3`

	return cart.DB.Exec(query, cartreq.Quantity, cartreq.CartID, cartreq.ProductId).Error

}

func (cart *CartAdapter) DeleteCartItem(cartreq helperstructs.CartItemReq) error {

	query := `DELETE FROM cart_items WHERE cart_id = $1 AND product_id = $2`

	return cart.DB.Exec(query, cartreq.CartID, cartreq.ProductId).Error

}

func (cart *CartAdapter) TruncateCart(cart_id uint) error {

	query := `DELETE FROM cart_items WHERE cart_id = $1`

	return cart.DB.Exec(query, cart_id).Error

}
