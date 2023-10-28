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

	query := `INSERT INTO carts (email,created_at) VALUES($1,NOW()) RETURNING id,email,created_at`

	return cart.DB.Raw(query, email).Error

}

func (cart *CartAdapter) AddToCart(cartreq helperstructs.CartItemReq) error {

	query := `INSERT INTO cart_items (cart_id,product_id,quantity,added_at) VALUES($1,$2,$3,NOW())`

	return cart.DB.Raw(query, cartreq.CartID, cartreq.ProductId, cartreq.Quantity).Error

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

func (cart *CartAdapter) GetItemByProductID(id uint) (responce.CartItemData, error) {

	var item responce.CartItemData

	query := `SELECT * FROM cart_items WHERE product_id = $1`

	return item, cart.DB.Raw(query, id).Scan(&item).Error

}
