package usecasesinterface

import (
	"context"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type AdminUsecaseInterface interface {
	AdminLogin(ctx context.Context, adminreq helperstructs.AdminReq) (responce.AdminData, error)
}