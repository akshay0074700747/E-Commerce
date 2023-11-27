package usecasesinterface

import (
	"context"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"time"
)

type AdminUsecaseInterface interface {
	AdminLogin(ctx context.Context, adminreq helperstructs.AdminReq) (responce.AdminData, error)
	GetUsers(ctx context.Context,count,page string) ([]responce.AdminsideUsersData, error)
	Reportuser(ctx context.Context, reportreq helperstructs.ReportReq) error
	GetUser(ctx context.Context, email string) (responce.AdminsideUsersData, error)
	GetAdminDashBoard(ctx context.Context) (responce.AdminDashBoard, error)
	GetUsersWalletDetails(ctx context.Context,count,page string) ([]responce.WalletsInfo,error)
	GetSalesReport(ctx context.Context, starttime,endtime time.Time) (responce.AdminSalesReport, error)
}
