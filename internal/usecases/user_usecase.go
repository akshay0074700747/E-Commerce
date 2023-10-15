package usecases

import (
	"ecommerce/internal/entities"
	"ecommerce/internal/interfaces/repositories"
	"fmt"
	"math/rand"
	"time"
)

type UserUsecase struct {
	user_repo repositories.UserRepo
}

func NewUserUsecase(user_repo repositories.UserRepo) *UserUsecase {

	return &UserUsecase{user_repo: user_repo}

}

func (user_usecase *UserUsecase) RegisterUser(email, password, mobile, name string) error {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	id := fmt.Sprintf("%06d", r.Intn(1000000))

	user := entities.User{
		Id:       id,
		Email:    email,
		Password: password,
		Mobile:   mobile,
		Name:     name,
	}

	return user_usecase.user_repo.CreateUser(&user)

}
