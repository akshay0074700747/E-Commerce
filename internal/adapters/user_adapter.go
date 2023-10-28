package adapters

import (
	"ecommerce/internal/interfaces/repositories"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) repositories.UserRepo {
	return &userDatabase{DB}
}

func (c *userDatabase) UserSignUp(user helperstructs.UserReq) (responce.UserData, error) {

	var userData responce.UserData

	insertQuery := `INSERT INTO users (name,email,mobile,password,created_at)VALUES($1,$2,$3,$4,$5) 
					RETURNING id,name,email,mobile`

	err := c.DB.Raw(insertQuery, user.Name, user.Email, user.Mobile, user.Password, time.Now()).Scan(&userData).Error

	return userData, err
}

func (c *userDatabase) GetByEmail(user helperstructs.UserReq) (responce.UserData, error) {

	var userdta responce.UserData

	selectquery := `SELECT * FROM users WHERE email = $1`

	err := c.DB.Raw(selectquery, user.Email).Scan(&userdta).Error

	return userdta, err

}

func (c *userDatabase) ReportAdmin(reportreq helperstructs.ReportReq) (error) {
	
	Insertquery := `INSERT INTO reports (email,description) VALUES ($1,$2)`

	result := c.DB.Exec(Insertquery, reportreq.Email, reportreq.Description)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows were affected by the insertion")
	}

	return nil

}
