package adapters

import (
	"ecommerce/internal/interfaces/repositories"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type AdminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) repositories.AdminRepo {

	return &AdminDatabase{DB: db}

}

func (admn *AdminDatabase) GetByEmail(admin helperstructs.AdminReq) (responce.AdminData, error) {

	var admindta responce.AdminData

	selectquery := `SELECT * FROM admins WHERE email = $1`

	err := admn.DB.Raw(selectquery, admin.Email).Scan(&admindta).Error

	return admindta, err

}

func (admn *AdminDatabase) GetAllUsers() ([]responce.AdminsideUsersData, error) {

	var usersdata []responce.AdminsideUsersData

	selectquery := `SELECT * FROM users`

	err := admn.DB.Raw(selectquery).Scan(&usersdata).Error

	return usersdata, err
}

func (admn *AdminDatabase) ReportUser(reportreq helperstructs.ReportReq) error {

	Insertquery := `INSERT INTO reports (email,description) VALUES ($1,$2)`

	result := admn.DB.Exec(Insertquery, reportreq.Email, reportreq.Description)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows were affected by the insertion")
	}

	return nil
}

func (admn *AdminDatabase) GetReports(email string) (int, error) {

	var count int

	query := `SELECT COUNT(email) FROM reports WHERE email = $1`

	return count, admn.DB.Raw(query, email).Scan(&count).Error

}

func (admn *AdminDatabase) GetUser(email string) (responce.AdminsideUsersData, error) {

	var userdata responce.AdminsideUsersData

	selectquery := `SELECT * FROM users WHERE email = $1`

	query := `SELECT COUNT(email) FROM reports WHERE email = $1`

	if err := admn.DB.Raw(selectquery, email).Scan(&userdata).Error; err != nil {
		return userdata, err
	}

	return userdata, admn.DB.Raw(query, email).Scan(&userdata.Reports).Error

}

func (admn *AdminDatabase) GetCountOfUsers() (int, error) {

	var count int

	query := `SELECT COUNT(*) FROM users;`

	return count, admn.DB.Raw(query).Scan(&count).Error

}

func (admn *AdminDatabase) GetCountOfProducts() (int, error) {

	var count int

	query := `SELECT COUNT(*) FROM products;`

	return count, admn.DB.Raw(query).Scan(&count).Error

}

func (admn *AdminDatabase) GetTotalSales() (int, error) {

	var count int

	query := `SELECT COUNT(*) FROM orders WHERE is_cancelled = false AND return_status = false;`

	return count, admn.DB.Raw(query).Scan(&count).Error

}

func (admn *AdminDatabase) GetTotalCancelledOrders() (int, error) {

	var count int

	query := `SELECT COUNT(*) FROM orders WHERE is_cancelled = true;`

	return count, admn.DB.Raw(query).Scan(&count).Error

}

func (admn *AdminDatabase) GetDeliveredOrders() (int, error) {

	var count int

	query := `SELECT COUNT(*) FROM orders WHERE status = 'delivered';`

	return count, admn.DB.Raw(query).Scan(&count).Error

}

func (admn *AdminDatabase) GetPurchasedUsers() (int, error) {

	var count int

	query := `SELECT COUNT(DISTINCT email) FROM orders WHERE is_cancelled = false AND return_status = false;`

	return count, admn.DB.Raw(query).Scan(&count).Error

}

func (admn *AdminDatabase) ActiveDiscounts() ([]uint, error) {

	var ids []uint

	query := `SELECT id FROM discounts;`

	return ids, admn.DB.Raw(query).Scan(&ids).Error

}

func (admn *AdminDatabase) TotalBlockedUsers() (int, error) {

	var count int

	query := `SELECT COUNT(*) FROM users WHERE isblocked = true;`

	return count, admn.DB.Raw(query).Scan(&count).Error

}

func (admn *AdminDatabase) BestSellerProduct() (string, error) {

	var name string

	query := `SELECT p.name FROM products p JOIN ( SELECT product_id FROM order_items GROUP BY product_id ORDER BY COUNT(*) DESC LIMIT 1 ) o ON p.id = o.product_id;`

	return name, admn.DB.Raw(query).Scan(&name).Error

}

func (admn *AdminDatabase) GetAllOrderedProductIDs() ([]responce.OrderedProductWithCount, error) {

	var prod []responce.OrderedProductWithCount

	query := `SELECT product_id,COUNT(*) AS count FROM order_items GROUP BY product_id`

	return prod, admn.DB.Raw(query).Scan(&prod).Error

}

func (admn *AdminDatabase) GetCategoryByProductID(prodid uint) (responce.CategoryData, error) {

	var cat responce.CategoryData

	query := `SELECT * FROM categories WHERE id = (SELECT category_id FROM products WHERE id = $1)`

	return cat, admn.DB.Raw(query, prodid).Scan(&cat).Error

}

func (admn *AdminDatabase) GetCategoryIDbyProdID(prodid uint) (uint, error) {

	var cat uint

	query := `SELECT category FROM products WHERE id = $1`

	return cat, admn.DB.Raw(query, prodid).Scan(&cat).Error

}

func (admn *AdminDatabase) GetCategoryByCatID(catid uint) (responce.CategoryData, error) {

	var cat responce.CategoryData

	query := `SELECT * FROM categories WHERE id = $1`

	return cat, admn.DB.Raw(query, catid).Scan(&cat).Error

}

func (admn *AdminDatabase) GetOrdrsByTime(timee time.Time) (int, error) {

	var count int

	query := `SELECT COUNT(*) FROM orders WHERE order_date >= $1 AND is_cancelled = false AND return_status = false`

	return count, admn.DB.Raw(query, timee).Scan(&count).Error

}

func (admn *AdminDatabase) GetMoneyByTime(timee time.Time) (int, error) {

	var count int

	query := `SELECT SUM(price) FROM orders WHERE order_date >= $1 AND is_cancelled = false AND return_status = false`

	admn.DB.Raw(query, timee).Scan(&count)

	return count, nil

}

func (admn *AdminDatabase) GetProductsSoldByTime(timee time.Time) (int, error) {

	var count int

	query := `SELECT COUNT(DISTINCT product_id) FROM order_items WHERE order_id IN (SELECT id FROM orders WHERE order_date >= $1)`

	return count, admn.DB.Raw(query, timee).Scan(&count).Error

}

func (admn *AdminDatabase) GetUsersOrderedByTime(timee time.Time) (int, error) {

	var count int

	query := `SELECT COUNT(DISTINCT email) FROM orders WHERE order_date >= $1 AND is_cancelled = false AND return_status = false;`

	return count, admn.DB.Raw(query, timee).Scan(&count).Error

}
