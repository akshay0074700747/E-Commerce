package usecasesinterface

import (
	"context"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type SuAdminUsecaseInterface interface {
	SuAdminLogin(ctx context.Context, suadmin helperstructs.SuAdminReq) (responce.SuAdminData, error)
	CreateAdmin(ctx context.Context, admin helperstructs.AdminReq) (responce.AdminData, error)
}