package adapters

import (
	"ecommerce/internal/interfaces/repositories"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"

	"gorm.io/gorm"
)

type WishListAdapter struct {
	DB *gorm.DB
}

func NewWishListAdapter(db *gorm.DB) repositories.WishListRepo {

	return &WishListAdapter{DB: db}

}

func (wishlist *WishListAdapter) CreateWishList(email string) error {

	query := `INSERT INTO wish_lists (email) VALUES($1)`

	return wishlist.DB.Exec(query, email).Error

}

func (wishlist *WishListAdapter) AddToWishList(wishreq helperstructs.WishListItemsReq) error {

	query := `INSERT INTO wish_list_items (wish_list_id,product_id) VALUES($1,$2)`

	return wishlist.DB.Exec(query, wishreq.WishListID, wishreq.ProductId).Error

}

func (wishlist *WishListAdapter) GetWishListItems(id uint) ([]responce.WishListItemsData, error) {

	var itemdata []responce.WishListItemsData

	query := `SELECT * FROM wish_list_items WHERE wish_list_id = $1`

	return itemdata, wishlist.DB.Raw(query, id).Scan(&itemdata).Error

}

func (wishlist *WishListAdapter) GetWishListID(email string) (uint, error) {

	var id uint

	query := `SELECT id FROM wish_lists WHERE email = $1`

	return id, wishlist.DB.Raw(query, email).Scan(&id).Error

}

func (wishlist *WishListAdapter) GetProductByID(id uint) (responce.ProuctData, error) {

	var products responce.ProuctData

	query := `SELECT * FROM products WHERE id = $1`

	return products, wishlist.DB.Raw(query, id).Scan(&products).Error

}

func (wishlist *WishListAdapter) GetItemByProductID(wishlist_id, product_id uint) (responce.WishListItemsData, error) {

	var item responce.WishListItemsData

	query := `SELECT * FROM wish_list_items WHERE wish_list_id = $1 AND product_id = $2`

	return item, wishlist.DB.Raw(query, wishlist_id, product_id).Scan(&item).Error

}

func (wishlist *WishListAdapter) DeleteWishListItem(wishlistreq helperstructs.WishListItemsReq) error {

	query := `DELETE FROM wish_list_items WHERE wish_list_id = $1 AND product_id = $2`

	return wishlist.DB.Exec(query, wishlistreq.WishListID, wishlistreq.ProductId).Error

}

func (wishlist *WishListAdapter) TruncateCart(wishlist_id uint) ([]responce.WishListItemsData, error) {

	var items []responce.WishListItemsData

	query := `DELETE FROM wish_list_items WHERE wish_list_id = $1 RETURNING product_id`

	return items, wishlist.DB.Raw(query, wishlist_id).Scan(&items).Error

}
