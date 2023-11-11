package repositories

import (
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type ReviewRepo interface {
	CreateReview(helperstructs.ReviewReq) error
	UpdateReview(helperstructs.ReviewReq) error
	GetReviewsByID(prodid uint) ([]responce.ReviewResponce, error)
	GetReviwByEmail(email string) ([]responce.ReviewResponce, error)
	DeleteReview(id uint) error
	VerifyOrderedUser(email string,productid uint) (error)
}
