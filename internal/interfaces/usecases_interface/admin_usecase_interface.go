package usecasesinterface

import (
	"context"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"time"
)

type AdminUsecaseInterface interface {
	AdminLogin(ctx context.Context, adminreq helperstructs.AdminReq) (responce.AdminData, error)
	GetUsers(ctx context.Context) ([]responce.AdminsideUsersData, error)
	Reportuser(ctx context.Context, reportreq helperstructs.ReportReq) error
	GetUser(ctx context.Context, email string) (responce.AdminsideUsersData, error)
	GetAdminDashBoard(ctx context.Context) (responce.AdminDashBoard, error)
	GetSalesReport(ctx context.Context, timee time.Time) (responce.AdminSalesReport, error)
}
