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

func (admn *AdminDatabase) ReportUser(reportreq helperstructs.ReportReq) error {

	Insertquery := `INSERT INTO reports (email,description) VALUES ($1,$2)`

	result := admn.DB.Exec(Insertquery, reportreq.Email, reportreq.Description)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows were affected by the insertion")
	}

	return nil
}

func (admn *AdminDatabase) GetReports(email string) (int, error) {

	var count int

	query := `SELECT COUNT(email) FROM reports WHERE email = $1`

	return count, admn.DB.Raw(query, email).Scan(&count).Error

}

func (admn *AdminDatabase) GetUser(email string) (responce.AdminsideUsersData, error) {

	var userdata responce.AdminsideUsersData

	selectquery := `SELECT * FROM users WHERE email = $1`

	query := `SELECT COUNT(email) FROM reports WHERE email = $1`

	if err := admn.DB.Raw(selectquery, email).Scan(&userdata).Error; err != nil {
		return userdata, err
	}

	return userdata, admn.DB.Raw(query, email).Scan(&userdata.Reports).Error

}
