package repositories

import (
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"time"
)

type AdminRepo interface {
	GetByEmail(admin helperstructs.AdminReq) (responce.AdminData, error)
	GetAllUsers() ([]responce.AdminsideUsersData, error)
	ReportUser(reportreq helperstructs.ReportReq) error
	GetReports(email string) (int, error)
	GetUser(email string) (responce.AdminsideUsersData, error)
	GetCountOfUsers() (int, error)
	GetCountOfProducts()(int,error)
	GetTotalSales()(int,error)
	GetTotalCancelledOrders()(int,error)
	GetDeliveredOrders()(int,error)
	GetPurchasedUsers()(int,error)
	ActiveDiscounts()([]uint,error)
	TotalBlockedUsers()(int,error)
	BestSellerProduct()(string,error)
	GetAllOrderedProductIDs()([]responce.OrderedProductWithCount,error)
	GetCategoryByProductID(prodid uint)(responce.CategoryData,error)
	GetCategoryIDbyProdID(prodid uint) (uint,error)
	GetCategoryByCatID(catid uint) (responce.CategoryData,error)
	GetOrdrsByTime(time.Time) (int,error)
	GetMoneyByTime(time.Time) (int,error)
	GetProductsSoldByTime(time.Time) (int,error)
	GetUsersOrderedByTime(time.Time) (int,error)
}