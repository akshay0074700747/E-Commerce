package helperstructs

type WishListsReq struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

type WishListItemsReq struct {
	ID         uint   `json:"id"`
	WishListID uint   `json:"wish_list_id"`
	ProductId  uint   `json:"product_id"`
	Email      string `json:"email"`
}
