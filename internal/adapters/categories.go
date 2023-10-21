package adapters

import (
	"ecommerce/internal/interfaces/repositories"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"

	"gorm.io/gorm"
)

type CategoryDataBase struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) repositories.CategoryRepo {

	return &CategoryDataBase{DB: db}

}

func (cat *CategoryDataBase) CreateCategory(catreq helperstructs.CategoryReq) (responce.CategoryData, error) {

	var catdata responce.CategoryData

	insertquery := `INSERT INTO categories (category,sub_category,update_by) VALUES ($1,$2,$3)
	 RETURNING id,category,sub_category,updated_by`

	err := cat.DB.Raw(insertquery, catreq.Category, catreq.SubCategory, catreq.UpdatedBy).Scan(&catdata).Error

	return catdata, err

}

func (cat *CategoryDataBase) UpdateCategory(catreq helperstructs.CategoryReq) (responce.CategoryData, error) {

	var catdata responce.CategoryData

	updatequery := `UPDATE categories SET category = $1, sub_category = $2, updated_by = $3 WHERE id = $4`

	err := cat.DB.Raw(updatequery, catreq.Category, catreq.SubCategory, catreq.UpdatedBy, catreq.Id).Scan(&catdata).Error

	return catdata, err

}

func (cat *CategoryDataBase) DeleteCategory(category_id uint) error {

	deletequery := `DELETE FROM categories WHERE id = $1`

	return cat.DB.Raw(deletequery, category_id).Error

}

func (cat *CategoryDataBase) GetallCategory() ([]responce.CategoryData, error) {

	var catdata []responce.CategoryData

	selectquesy := `SELECT * FROM categories`

	err := cat.DB.Raw(selectquesy).Scan(&catdata).Error

	return catdata, err

}
