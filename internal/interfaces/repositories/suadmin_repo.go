package repositories

import (
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type SuAdminRepo interface {
	GetByEmail(suadmin helperstructs.SuAdminReq) (responce.SuAdminData,error)
	CreateAdmin(admin helperstructs.AdminReq) (responce.AdminData, error)
}