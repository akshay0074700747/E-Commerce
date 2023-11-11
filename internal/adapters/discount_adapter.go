package adapters

import (
	"ecommerce/internal/interfaces/repositories"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"

	"gorm.io/gorm"
)

type DiscountAdapter struct {
	DB *gorm.DB
}

func NewDiscountAdapter(db *gorm.DB) repositories.DiscountRepo {
	return &DiscountAdapter{DB: db}
}

func (discount *DiscountAdapter) AddDiscount(discountreq helperstructs.DiscountReq) (responce.DiscountData, error) {

	var discountdta responce.DiscountData

	insertquery := `INSERT INTO discounts (category,discount,end_date) VALUES ($1,$2,$3) RETURNING id,category,discount,end_date`

	return discountdta, discount.DB.Raw(insertquery, discountreq.Category, discountreq.Discount, discountreq.EndDate).Scan(&discountdta).Error

}

func (discount *DiscountAdapter) UpdateDiscount(discountreq helperstructs.DiscountReq) (responce.DiscountData, error) {

	var discountdta responce.DiscountData

	updatequery := `UPDATE discounts SET category = $1, discount = $2, end_date = $3 WHERE id = $4 RETURNING id,category,discount,end_date`

	return discountdta, discount.DB.Raw(updatequery, discountreq.Category, discountreq.Discount, discountreq.EndDate, discountreq.ID).Scan(&discountdta).Error

}

func (discount *DiscountAdapter) DeleteDiscount(id uint) error {

	deletequery := `DELETE FROM discounts WHERE id = $1`

	return discount.DB.Exec(deletequery, id).Error

}

func (discount *DiscountAdapter) GetAllDiscounts() ([]responce.DiscountData, error) {

	selectquery := `SELECT * FROM discounts`

	var discountdata []responce.DiscountData

	return discountdata, discount.DB.Raw(selectquery).Scan(&discountdata).Error

}

func (discount *DiscountAdapter) GetByID(category_id uint) (responce.DiscountData, error) {

	var discountdata responce.DiscountData

	selectquery := `SELECT * FROM discounts WHERE category = $1`

	discount.DB.Raw(selectquery, category_id).Scan(&discountdata)

	return discountdata, nil

}

func (discount *DiscountAdapter) GetByProductID(prod_id uint) (responce.DiscountData, error) {

	var discountdata responce.DiscountData

	selectquery := `SELECT * FROM discounts WHERE category = (SELECT category FROM products WHERE id = $1)`

	discount.DB.Raw(selectquery, prod_id).Scan(&discountdata)

	return discountdata, nil

}
