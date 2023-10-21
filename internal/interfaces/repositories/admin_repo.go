package repositories

import (
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type AdminRepo interface {
	GetByEmail(admin helperstructs.AdminReq) (responce.AdminData, error)
	GetAllUsers() ([]responce.AdminsideUsersData, error)
	ReportUser(email string) (error)
}
