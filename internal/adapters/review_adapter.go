package adapters

import (
	"ecommerce/internal/interfaces/repositories"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"fmt"

	"gorm.io/gorm"
)

type ReviewAdapter struct {
	DB *gorm.DB
}

func NewReviewAdapter(db *gorm.DB) repositories.ReviewRepo {
	return &ReviewAdapter{
		DB: db,
	}
}

func (review *ReviewAdapter) CreateReview(revreq helperstructs.ReviewReq) error {

	query := `INSERT INTO reviews (product,reviewed_by,description,rating) VALUES($1,$2,$3,$4)`

	return review.DB.Exec(query, revreq.Product, revreq.ReviewedBy, revreq.Description, revreq.Rating).Error

}

func (review *ReviewAdapter) UpdateReview(revreq helperstructs.ReviewReq) error {

	query := `UPDATE reviews SET product = $1,reviewed_by = $2,description = $3,rating = $4`

	return review.DB.Exec(query, revreq.Product, revreq.ReviewedBy, revreq.Description, revreq.Rating).Error

}

func (review *ReviewAdapter) GetReviewsByID(prodid uint) ([]responce.ReviewResponce, error) {

	var revres []responce.ReviewResponce

	query := `SELECT * FROM reviews WHERE product = $1`

	return revres, review.DB.Raw(query, prodid).Scan(&revres).Error

}

func (review *ReviewAdapter) GetReviwByEmail(email string) ([]responce.ReviewResponce, error) {

	var revres []responce.ReviewResponce

	query := `SELECT * FROM reviews WHERE reviewed_by = $1`

	return revres, review.DB.Raw(query, email).Scan(&revres).Error

}

func (review *ReviewAdapter) DeleteReview(id uint) error {

	query := `DELETE FROM reviews WHERE id = $1`

	return review.DB.Exec(query, id).Error

}

func (review *ReviewAdapter) VerifyOrderedUser(email string, productid uint) error {

	query := `SELECT * FROM order_items WHERE product_id = $1 AND order_id IN (SELECT id FROM orders WHERE email = $2)`

	result := review.DB.Exec(query, productid, email)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("you havent purchased this item yet to review it")
	}

	return nil

}
