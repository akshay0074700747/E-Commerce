package repositories

import (
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type AdminRepo interface {
	GetByEmail(admin helperstructs.AdminReq) (responce.AdminData, error)
	GetAllUsers() ([]responce.AdminsideUsersData, error)
	ReportUser(reportreq helperstructs.ReportReq) (error)
	GetReports(email string)(int,error)
	GetUser(email string) (responce.AdminsideUsersData,error)
}
