package adapters

import (
	"ecommerce/internal/interfaces/repositories"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"fmt"

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
		relative_products,updated_by,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,NOW()) RETURNING 
		id,category,brand,name,description,price,stock,specifications,
		relative_products,updated_by,updated_at`

	return productdta, product.DB.Raw(insertquery, productreq.Category, productreq.Brand, productreq.Name, productreq.Description,
		productreq.Price, productreq.Stock, productreq.Specifications,
		productreq.RelatedProducts, productreq.UpdatedBy).Scan(&productdta).Error

}

func (product *ProductDataBase) GetProducts(offset, count int) ([]responce.ProuctData, error) {

	var productdta []responce.ProuctData

	selectquery := `SELECT products.*,categories.category,categories.sub_category FROM products
	INNER JOIN categories ON products.category = categories.id OFFSET $1 LIMIT $2`

	if err := product.DB.Raw(selectquery, offset, count).Scan(&productdta).Error; err != nil {
		return productdta, err
	}

	querryyy := `SELECT relative_products from products WHERE id = $1`

	for i := range productdta {
		if err := product.DB.Raw(querryyy, productdta[i].ID).Scan(&productdta[i].RelativeProducts).Error; err != nil {
			fmt.Println(err.Error())
		}
	}

	return productdta, nil

}

func (product *ProductDataBase) UpdateProduct(productreq helperstructs.ProductReq) (responce.ProuctData, error) {
	var productdta responce.ProuctData

	updatequery := `UPDATE products SET category = $1, brand = $2, name = $3, description = $4, 
        price = $5, stock = $6, specifications = $7, relative_products = $8, updated_by = $9 
        WHERE id = $10 RETURNING id,category,brand,name,description,price,stock,specifications,
		relative_products,updated_by,updated_at`

	return productdta, product.DB.Raw(updatequery, productreq.Category, productreq.Brand, productreq.Name,
		productreq.Description, productreq.Price, productreq.Stock, productreq.Specifications,
		productreq.RelatedProducts, productreq.UpdatedBy, productreq.ID).Scan(&productdta).Error
}

func (product *ProductDataBase) DeleteProduct(product_id uint) error {

	deletequery := `DELETE FROM products WHERE id = $1`

	err := product.DB.Exec(deletequery, product_id).Error

	if err != nil {
		return err
	}

	deletequery = `DELETE FROM images WHERE product_id = $1`

	err = product.DB.Exec(deletequery, product_id).Error

	if err != nil {
		return err
	}

	return nil

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

func (product *ProductDataBase) UpdateStock(id uint, stock int) error {

	query := `UPDATE products SET stock = $1 WHERE id = $2`

	return product.DB.Exec(query, stock, id).Error

}

func (product *ProductDataBase) GetProductByID(id uint) (responce.ProuctData, error) {

	var products responce.ProuctData

	query := `SELECT * FROM products WHERE id = $1`

	return products, product.DB.Raw(query, id).Scan(&products).Error

}

func (product *ProductDataBase) GetPriceByID(id uint) (int, error) {

	var price int

	query := `SELECT price FROM products WHERE id = $1`

	return price, product.DB.Raw(query, id).Scan(&price).Error

}

func (product *ProductDataBase) IncreaseStock(id uint, stock int) error {

	query := `UPDATE products SET stock = stock + $1 WHERE id = $2`

	return product.DB.Exec(query, stock, id).Error

}

func (product *ProductDataBase) DecreaseStock(id uint, stock int) error {

	query := `UPDATE products SET stock = stock - $1 WHERE id = $2 AND stock >= $1 RETURNING stock`

	result := product.DB.Exec(query, stock, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("insufficient stock for product with id %d", id)
	}

	return nil

}

func (product *ProductDataBase) GetRatingByID(id uint) (float32, error) {

	var rating float32

	query := `SELECT AVG(rating) FROM reviews WHERE product = $1`

	product.DB.Raw(query, id).Scan(&rating)

	return rating, nil

}

func (product *ProductDataBase) AddImages(prod_id uint, image string) error {

	query := `INSERT INTO images (product_id,image) VALUES($1,$2);`

	return product.DB.Exec(query, prod_id, image).Error

}

func (product *ProductDataBase) GetAllImages(prod_id uint) ([]string, error) {

	var images []string

	query := `SELECT image FROM images WHERE product_id = $1`

	product.DB.Raw(query, prod_id).Scan(&images)

	return images, nil

}
