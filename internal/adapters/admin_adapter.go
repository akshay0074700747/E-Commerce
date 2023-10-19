package adapters

import (
	"ecommerce/internal/interfaces/repositories"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"

	"gorm.io/gorm"
)

type AdminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) repositories.AdminRepo {
	
	return &AdminDatabase{DB: db}

}

func (admn *AdminDatabase) GetByEmail(admin helperstructs.AdminReq) (responce.AdminData,error)  {
	
	var admindta responce.AdminData

	selectquery := `SELECT * FROM admins WHERE email = $1`

	err := admn.DB.Raw(selectquery,admin.Email).Scan(&admindta).Error

	return admindta,err

}
