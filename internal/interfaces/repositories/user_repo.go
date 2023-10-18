package repositories

import (
	"context"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type UserRepo interface {
	UserSignUp(ctx context.Context, user helperstructs.UserReq) (responce.UserData, error)
}
