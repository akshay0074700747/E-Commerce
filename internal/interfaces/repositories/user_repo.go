package repositories

import "ecommerce/internal/entities"

type UserRepo interface {
	CreateUser(user *entities.User) error
}