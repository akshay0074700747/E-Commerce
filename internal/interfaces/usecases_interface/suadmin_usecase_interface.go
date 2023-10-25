package usecasesinterface

import (
	"context"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type SuAdminUsecaseInterface interface {
	SuAdminLogin(ctx context.Context, suadmin helperstructs.SuAdminReq) (responce.SuAdminData, error)
	CreateAdmin(ctx context.Context, admin helperstructs.AdminReq) (responce.AdminData, error)
	BlockUser(ctx context.Context, blockreq helperstructs.BlockReq) error
	GetAllUsers(ctx context.Context) ([]responce.AdminsideUsersData, error)
	GetAllAdmins(ctx context.Context) ([]responce.AdminData, error)
	GetReportes(ctx context.Context) ([]responce.DetailReportResponce, error)
	GetDetailedReport(ctx context.Context, email string) (responce.DetailReportResponce, error)
}
