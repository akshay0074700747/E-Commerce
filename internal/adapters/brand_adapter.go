package adapters

import (
	"ecommerce/internal/interfaces/repositories"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"

	"gorm.io/gorm"
)

type BrandDataBase struct {
	DB *gorm.DB
}

func NewBrandRepository(db *gorm.DB) repositories.BrandRepo {

	return &BrandDataBase{DB: db}

}

func (brand *BrandDataBase) CreateBrand(brandreq helperstructs.BrandReq) (responce.BrandData, error) {

	var branddata responce.BrandData

	insertquery := `INSERT INTO brands (name) VALUES ($1)
	 RETURNING id,name`

	err := brand.DB.Raw(insertquery, brandreq.Name).Scan(&branddata).Error

	return branddata, err

}

func (brand *BrandDataBase) UpdateBrand(brandreq helperstructs.BrandReq) (responce.BrandData, error) {

	var branddata responce.BrandData

	updatequery := `UPDATE brands SET name = $1 WHERE id = $2 RETURNING id,name`

	err := brand.DB.Raw(updatequery, brandreq.Name, brandreq.ID).Scan(&branddata).Error

	return branddata, err

}

func (brand *BrandDataBase) DeleteBrand(brand_id uint) error {

	updatequery := `UPDATE products SET brand = NULL WHERE brand = $1;`

	deletequery := `DELETE FROM brands WHERE id = $1;`

	if err := brand.DB.Exec(updatequery, brand_id).Error; err != nil {
		return err
	}

	return brand.DB.Exec(deletequery, brand_id).Error

}

func (brand *BrandDataBase) GetallBrand() ([]responce.BrandData, error) {

	var branddata []responce.BrandData

	selectquesy := `SELECT * FROM brands`

	err := brand.DB.Raw(selectquesy).Scan(&branddata).Error

	return branddata, err

}
