package responce

type WishListsData struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

type WishListItemsData struct {
	ID        uint `json:"id"`
	WishListID    uint `json:"wish_list_id"`
	ProductId uint `json:"product_id"`
}
