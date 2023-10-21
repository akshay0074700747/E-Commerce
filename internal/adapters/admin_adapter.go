package adapters

import (
	"ecommerce/internal/interfaces/repositories"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"fmt"

	"gorm.io/gorm"
)

type AdminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) repositories.AdminRepo {

	return &AdminDatabase{DB: db}

}

func (admn *AdminDatabase) GetByEmail(admin helperstructs.AdminReq) (responce.AdminData, error) {

	var admindta responce.AdminData

	selectquery := `SELECT * FROM admins WHERE email = $1`

	err := admn.DB.Raw(selectquery, admin.Email).Scan(&admindta).Error

	return admindta, err

}

func (admn *AdminDatabase) GetAllUsers() ([]responce.AdminsideUsersData, error) {

	var usersdata []responce.AdminsideUsersData

	selectquery := `SELECT * FROM users`

	err := admn.DB.Raw(selectquery).Scan(&usersdata).Error

	return usersdata, err
}

func (admn *AdminDatabase) ReportUser(email string) error {

	updateQuery := `UPDATE users SET reports = reports + 1 WHERE email = ?`

	result := admn.DB.Exec(updateQuery, email)
	
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows were affected by the update")
	}

	return nil
}
