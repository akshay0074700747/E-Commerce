package adapters

import (
	"ecommerce/internal/interfaces/repositories"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"

	"gorm.io/gorm"
)

type ProductDataBase struct {
	DB *gorm.DB
}

func NewProductDataBase(repo *gorm.DB) repositories.ProductsRepo {

	return &ProductDataBase{DB: repo}

}

func (product *ProductDataBase) AddProduct(productreq helperstructs.ProductReq) (responce.ProuctData, error) {

	var productdta responce.ProuctData

	insertquery := `INSERT INTO products (category,brand,name,description,price,stock,specifications,
		is_active,relative_products,updated_by) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`

	return productdta, product.DB.Raw(insertquery, productreq.Category, productreq.Brand, productreq.Name, productreq.Description,
		productreq.Price, productreq.Stock, productreq.Specifications, productreq.IsActive,
		productreq.RelatedProducts, productreq.UpdatedBy).Scan(&productdta).Error

}

func (product *ProductDataBase) GetProducts() ([]responce.ProuctData, error) {

	var productdta []responce.ProuctData

	selectquery := `SELECT products.*,categories.category,categories.sub_category FROM products
	INNER JOIN categories ON products.category = categories.id`

	return productdta, product.DB.Raw(selectquery).Scan(&productdta).Error

}

func (product *ProductDataBase) UpdateProduct(productreq helperstructs.ProductReq) (responce.ProuctData, error) {
	var productdta responce.ProuctData

	updatequery := `UPDATE products SET category = $1, brand = $2, name = $3, description = $4, 
        price = $5, stock = $6, specifications = $7, is_active = $8, relative_products = $9, updated_by = $10 
        WHERE id = $11`

	return productdta, product.DB.Exec(updatequery, productreq.Category, productreq.Brand, productreq.Name,
		productreq.Description, productreq.Price, productreq.Stock, productreq.Specifications, productreq.IsActive,
		productreq.RelatedProducts, productreq.UpdatedBy, productreq.ID).Scan(&productdta).Error
}

func (product *ProductDataBase) DeleteProduct(product_id uint) error {

	deletequery := `DELETE FROM products WHERE id = $1`

	return product.DB.Exec(deletequery, product_id).Error

}

func (product *ProductDataBase) FindRelatedProducts(cat_id uint) ([]uint, error) {

	var id []uint

	query := `SELECT id FROM categories WHERE category = (SELECT category FROM categories WHERE id = $1)`

	return id, product.DB.Raw(query, cat_id).Scan(&id).Error

}

func (product *ProductDataBase) FindDiscountByID(category_id uint) (responce.DiscountData, error) {

	var discountdata responce.DiscountData

	selectquery := `SELECT * FROM discounts WHERE category = $1`

	return discountdata, product.DB.Raw(selectquery, category_id).Scan(&discountdata).Error

}

func (product *ProductDataBase) GetCategoryID(category, subcategory string) (uint, error) {

	var count uint

	query := `SELECT id FROM categories WHERE category = $1 AND sub_category = $2`

	return count, product.DB.Raw(query, category, subcategory).Scan(&count).Error

}

func (product *ProductDataBase) GetBrand(id uint) (string, error) {

	var brand string

	query := `SELECT name FROM brands WHERE id = $1`

	return brand, product.DB.Raw(query, id).Scan(&brand).Error

}
