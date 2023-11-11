package usecases

import (
	"context"
	"ecommerce/internal/interfaces/repositories"
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	"ecommerce/web/helpers"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"fmt"
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

	userData, err := c.userRepo.UserSignUp(user)

	return userData, err
}

func (c *userUseCase) UserLogin(ctx context.Context, user helperstructs.UserReq) (responce.UserData, error) {

	userdata, err := c.userRepo.GetByEmail(user)

	if err != nil {
		return responce.UserData{}, err
	}

	hashpass, err := c.userRepo.CheckPassword(user.Email)

	if err != nil {
		return responce.UserData{}, err
	}

	if err := helpers.VerifyPassword(user.Password, hashpass); err != nil {
		return responce.UserData{}, err
	}

	if userdata.Isblocked {
		return responce.UserData{}, fmt.Errorf("you have been blocked")
	}

	return userdata, nil

}

func (c *userUseCase) ReportAdmin(ctx context.Context, reportreq helperstructs.ReportReq) error {

	return c.userRepo.ReportAdmin(reportreq)

}

func (c *userUseCase) GetByEmail(ctx context.Context, user helperstructs.UserReq) (responce.UserData, error) {

	return c.userRepo.GetByEmail(user)

}

func (c *userUseCase) UpdateUserData(ctx context.Context, user helperstructs.UserReq) error {

	return c.userRepo.UpdateUserData(user)

}

func (c *userUseCase) ChangePassword(ctx context.Context, user helperstructs.UserReq) error {

	hashpass, err := c.userRepo.CheckPassword(user.Email)

	if err != nil {
		return err
	}

	if err := helpers.VerifyPassword(user.OldPassword, hashpass); err != nil {
		return err
	}

	hashh, err := helpers.Hash_pass(user.Password)

	if err != nil {
		return err
	}

	user.Password = string(hashh)

	return c.userRepo.ChangePassword(user)

}

func (c *userUseCase) ForgotPassword(ctx context.Context, user helperstructs.UserReq) error {

	hashh, err := helpers.Hash_pass(user.Password)

	if err != nil {
		return err
	}

	user.Password = string(hashh)

	return c.userRepo.ChangePassword(user)

}
