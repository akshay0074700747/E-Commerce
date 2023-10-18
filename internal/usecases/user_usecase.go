package usecases

import (
	"context"
	"ecommerce/internal/interfaces/repositories"
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	"ecommerce/web/helpers"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type userUseCase struct {
	userRepo repositories.UserRepo
}

func NewUserUseCase(repo repositories.UserRepo) usecasesinterface.UserUsecaseInterface {
	return &userUseCase{
		userRepo: repo,
	}
}

func (c *userUseCase) UserSignUp(ctx context.Context, user helperstructs.UserReq) (responce.UserData, error) {
	hash, err := helpers.Hash_pass(user.Password)
	if err != nil {
		return responce.UserData{}, err
	}
	user.Password = string(hash)
	userData, err := c.userRepo.UserSignUp(ctx, user)
	return userData, err
}
