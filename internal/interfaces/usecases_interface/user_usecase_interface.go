package usecasesinterface

import (
	"context"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type UserUsecaseInterface interface {
	UserSignUp(ctx context.Context, user helperstructs.UserReq) (responce.UserData, error)
	UserLogin(ctx context.Context, user helperstructs.UserReq) (responce.UserData, error)
	ReportAdmin(ctx context.Context, reportreq helperstructs.ReportReq) error
	GetByEmail(ctx context.Context, user helperstructs.UserReq) (responce.UserData, error)
	UpdateUserData(ctx context.Context, user helperstructs.UserReq) error
	ChangePassword(ctx context.Context, user helperstructs.UserReq) error
	ForgotPassword(ctx context.Context, user helperstructs.UserReq) error
}
