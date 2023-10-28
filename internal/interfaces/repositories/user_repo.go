package repositories

import (
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type UserRepo interface {
	UserSignUp(user helperstructs.UserReq) (responce.UserData, error)
	GetByEmail(user helperstructs.UserReq) (responce.UserData, error)
	ReportAdmin(reportreq helperstructs.ReportReq) (error)
}
