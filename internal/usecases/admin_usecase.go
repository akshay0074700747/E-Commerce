package usecases

import (
	"context"
	"ecommerce/internal/interfaces/repositories"
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"fmt"
	"time"
)

type AdminUsecase struct {
	AdminRepo repositories.AdminRepo
}

func NewAdminUsecase(repo repositories.AdminRepo) usecasesinterface.AdminUsecaseInterface {

	return &AdminUsecase{AdminRepo: repo}

}

func (admin *AdminUsecase) AdminLogin(ctx context.Context, adminreq helperstructs.AdminReq) (responce.AdminData, error) {

	admindata, err := admin.AdminRepo.GetByEmail(adminreq)

	if err != nil {
		return admindata, err
	}

	if admindata.Isblocked {
		return responce.AdminData{}, fmt.Errorf("sorry you have been blocked")
	}

	return admindata, nil

}

func (admin *AdminUsecase) GetUsers(ctx context.Context) ([]responce.AdminsideUsersData, error) {

	userdata, err := admin.AdminRepo.GetAllUsers()

	if err != nil {
		return userdata, err
	}

	for i := range userdata {

		count, err := admin.AdminRepo.GetReports(userdata[i].Email)

		if err != nil {
			return userdata, err
		}

		userdata[i].Reports = count
	}

	return userdata, nil

}

func (admin *AdminUsecase) Reportuser(ctx context.Context, reportreq helperstructs.ReportReq) error {

	return admin.AdminRepo.ReportUser(reportreq)

}

func (admin *AdminUsecase) GetUser(ctx context.Context, email string) (responce.AdminsideUsersData, error) {

	return admin.AdminRepo.GetUser(email)

}

func (admin *AdminUsecase) GetAdminDashBoard(ctx context.Context) (responce.AdminDashBoard, error) {

	var dashboard responce.AdminDashBoard
	var err error

	if dashboard.Users, err = admin.AdminRepo.GetCountOfUsers(); err != nil {
		return dashboard, err
	}

	if dashboard.Products, err = admin.AdminRepo.GetCountOfProducts(); err != nil {
		return dashboard, err
	}

	if dashboard.Sales, err = admin.AdminRepo.GetTotalSales(); err != nil {
		return dashboard, err
	}

	if dashboard.CancelledOrders, err = admin.AdminRepo.GetTotalCancelledOrders(); err != nil {
		return dashboard, err
	}

	if dashboard.DeliveredOrders, err = admin.AdminRepo.GetDeliveredOrders(); err != nil {
		return dashboard, err
	}

	if dashboard.PurchasedUsers, err = admin.AdminRepo.GetPurchasedUsers(); err != nil {
		return dashboard, err
	}

	if dashboard.ActiveDiscounts, err = admin.AdminRepo.ActiveDiscounts(); err != nil {
		return dashboard, err
	}

	if dashboard.TotalBlockedUsers, err = admin.AdminRepo.TotalBlockedUsers(); err != nil {
		return dashboard, err
	}

	if dashboard.MostPurchasedProduct, err = admin.AdminRepo.BestSellerProduct(); err != nil {
		return dashboard, err
	}

	var prods []responce.OrderedProductWithCount
	var catfinder = make(map[uint]int)
	var cat uint
	var largcat uint
	var category responce.CategoryData

	if prods, err = admin.AdminRepo.GetAllOrderedProductIDs(); err != nil {
		return dashboard, err
	}

	for _, prod := range prods {

		if cat, err = admin.AdminRepo.GetCategoryIDbyProdID(prod.ProductID); err != nil {
			return dashboard, err
		}

		catfinder[cat] += prod.Count

		if catfinder[largcat] < catfinder[cat] {
			largcat = cat
		}

	}

	if category, err = admin.AdminRepo.GetCategoryByCatID(largcat); err != nil {
		return dashboard, err
	}

	dashboard.MostPurchasedCategory = category.Category
	dashboard.MostPurchasedSubCategory = category.SubCategory

	return dashboard, nil

}

func (admin *AdminUsecase) GetSalesReport(ctx context.Context, timee time.Time) (responce.AdminSalesReport, error) {

	var salesreport responce.AdminSalesReport
	var err error

	if salesreport.Orders, err = admin.AdminRepo.GetOrdrsByTime(timee); err != nil {
		return salesreport, err
	}

	if salesreport.TransactionAmount, err = admin.AdminRepo.GetMoneyByTime(timee); err != nil {
		return salesreport, err
	}

	if salesreport.BuyedUsers, err = admin.AdminRepo.GetUsersOrderedByTime(timee); err != nil {
		return salesreport, err
	}

	if salesreport.ProductsSold, err = admin.AdminRepo.GetProductsSoldByTime(timee); err != nil {
		return salesreport, err
	}

	return salesreport, err

}
