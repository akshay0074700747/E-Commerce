package adapters

import (
	"ecommerce/internal/entities"

	"gorm.io/gorm"
)

type UserAdapter struct {
	db *gorm.DB
}

func NewUserAdapter(db *gorm.DB) *UserAdapter {

	return &UserAdapter{db: db}

}

func (useradapter *UserAdapter) CreateUser(user *entities.User) error {

	return useradapter.db.Create(user).Error

}