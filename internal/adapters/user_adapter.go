package adapters

import (
	"context"
	"ecommerce/internal/interfaces/repositories"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"

	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) repositories.UserRepo {
	return &userDatabase{DB}
}

func (c *userDatabase) UserSignUp(ctx context.Context, user helperstructs.UserReq) (responce.UserData, error) {
	var userData responce.UserData
	insertQuery := `INSERT INTO user (name,email,mobile,password)VALUES($1,$2,$3,$4) 
					RETURNING id,name,email,mobile`
	err := c.DB.Raw(insertQuery, user.Name, user.Email, user.Mobile, user.Password).Scan(&userData).Error
	return userData, err
}
