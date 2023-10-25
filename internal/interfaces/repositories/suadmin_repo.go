package repositories

import (
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type SuAdminRepo interface {
	GetByEmail(suadmin helperstructs.SuAdminReq) (responce.SuAdminData, error)
	CreateAdmin(admin helperstructs.AdminReq) (responce.AdminData, error)
	BlockUser(blockreq helperstructs.BlockReq) error
	GetReports(email string)(int,error)
	GetAllUsers() ([]responce.AdminsideUsersData, error)
	GetAllAdmins() ([]responce.AdminData, error)
	GetReportes() ([]responce.DetailReportResponce, error)
	GetDetailedReport(email string) (responce.DetailReportResponce, error)
}
