package usecasesinterface

import (
	"context"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type UserUsecaseInterface interface {
	UserSignUp(ctx context.Context, user helperstructs.UserReq) (responce.UserData, error)
}
