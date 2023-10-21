package usecasesinterface

import (
	"context"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type AdminUsecaseInterface interface {
	AdminLogin(ctx context.Context, adminreq helperstructs.AdminReq) (responce.AdminData, error)
	GetUsers(ctx context.Context) ([]responce.AdminsideUsersData,error)
	Reportuser(ctx context.Context, email string) (error)
}