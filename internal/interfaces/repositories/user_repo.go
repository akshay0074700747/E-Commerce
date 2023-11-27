package repositories

import (
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type UserRepo interface {
	UserSignUp(user helperstructs.UserReq) (responce.UserData, error)
	GetByEmail(user helperstructs.UserReq) (responce.UserData, error)
	ReportAdmin(reportreq helperstructs.ReportReq) error
	IncrementWallet(email string, money int) error
	DecrementWallet(email string, money int) error
	UpdateUserData(user helperstructs.UserReq) error
	CheckPassword(email string) (string, error)
	ChangePassword(user helperstructs.UserReq) error
	GetEmailByReferral(refid uint) (string,error)
	GetWalletByEmail(email string) (int,error)
}
