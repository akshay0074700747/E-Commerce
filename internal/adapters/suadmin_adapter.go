package adapters

import (
	"ecommerce/internal/interfaces/repositories"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"time"

	"gorm.io/gorm"
)

type SuAdminDataBase struct {
	DB *gorm.DB
}

func NewSuAdminRepository(db *gorm.DB) repositories.SuAdminRepo {
	return &SuAdminDataBase{DB: db}
}

func (suadmn *SuAdminDataBase) GetByEmail(suadmin helperstructs.SuAdminReq) (responce.SuAdminData, error) {

	var suadmndta responce.SuAdminData

	selectquery := `SELECT * FROM super_admins WHERE email = $1`

	err := suadmn.DB.Raw(selectquery, suadmin.Email).Scan(&suadmndta).Error

	return suadmndta, err
}

func (suadmn *SuAdminDataBase) CreateAdmin(admin helperstructs.AdminReq) (responce.AdminData, error) {

	var admindata responce.AdminData

	insertquery := `INSERT INTO admins (email,password,name,created_at) VALUES ($1,$2,$3,$4) RETURNING id,email,name`

	err := suadmn.DB.Raw(insertquery, admin.Email, admin.Password, admin.Name, time.Now()).Scan(&admindata).Error

	return admindata, err

}

func (suadmn *SuAdminDataBase) BlockUser(blockreq helperstructs.BlockReq) error {

	var updatequery string

	var checkadmn responce.AdminData

	checkadmnquery := `SELECT * FROM admins WHERE email = $1`

	suadmn.DB.Raw(checkadmnquery, blockreq.Email).Scan(&checkadmn)

	if checkadmn.Email != "" {
		updatequery = `UPDATE admins SET isblocked = NOT isblocked,unblock_time = $1 WHERE email = $2`
	} else {
		updatequery = `UPDATE users SET isblocked = NOT isblocked,unblock_time = $1 WHERE email = $2`
	}

	deletequery := `DELETE FROM reports WHERE email = $1`

	if err := suadmn.DB.Exec(updatequery, blockreq.Time, blockreq.Email).Error; err != nil {
		return err
	}

	return suadmn.DB.Exec(deletequery, blockreq.Email).Error

}

func (suadmin *SuAdminDataBase) GetAllUsers() ([]responce.AdminsideUsersData, error) {

	var usersdata []responce.AdminsideUsersData

	selectquery := `SELECT * FROM users`

	err := suadmin.DB.Raw(selectquery).Scan(&usersdata).Error

	return usersdata, err

}

func (suadmin *SuAdminDataBase) GetAllAdmins() ([]responce.AdminData, error) {

	var admindta []responce.AdminData

	selectquery := `SELECT * FROM admins`

	err := suadmin.DB.Raw(selectquery).Scan(&admindta).Error

	return admindta, err

}

func (suadmin *SuAdminDataBase) GetReportes() ([]responce.DetailReportResponce, error) {

	var reports []responce.DetailReportResponce

	query := `SELECT email,COUNT(*) AS reports,ARRAY_AGG(description) AS description FROM reports GROUP BY email`

	return reports, suadmin.DB.Raw(query).Scan(&reports).Error

}

func (suadmin *SuAdminDataBase) GetDetailedReport(email string) (responce.DetailReportResponce, error) {

	var report responce.DetailReportResponce

	query := `SELECT email,COUNT(*) AS reports,ARRAY_AGG(description) AS description FROM reports WHERE email = $1 GROUP BY email`

	return report, suadmin.DB.Raw(query, email).Scan(&report).Error

}

func (suadmin *SuAdminDataBase) GetReports(email string) (int, error) {

	var count int

	query := `SELECT COUNT(email) FROM reports WHERE email = $1`

	return count, suadmin.DB.Raw(query, email).Scan(&count).Error

}
