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

	selectquery := `SELECT * FROM superadmins WHERE email = $1`

	err := suadmn.DB.Raw(selectquery, suadmin.Email).Scan(&suadmndta).Error

	return suadmndta, err
}

func (suadmn *SuAdminDataBase) CreateAdmin(admin helperstructs.AdminReq) (responce.AdminData, error) {

	var admindata responce.AdminData

	insertquery := `INSERT INTO admins (email,password,name,created_at) VALUES ($1,$2,$3,$4) RETURNING id,email,name`

	err := suadmn.DB.Raw(insertquery, admin.Email, admin.Password, admin.Name, time.Now()).Scan(&admindata).Error

	return admindata, err

}
