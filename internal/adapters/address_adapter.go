package adapters

import (
	"ecommerce/internal/interfaces/repositories"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"

	"gorm.io/gorm"
)

type AddressAdapter struct {
	DB *gorm.DB
}

func NewAddressAdapter(db *gorm.DB) repositories.AddressRepo {
	return &AddressAdapter{DB: db}
}

func (address *AddressAdapter) GetallAddress(email string) ([]responce.AddressData, error) {

	var addressres []responce.AddressData

	query := `SELECT * FROM addresses where email = $1`

	return addressres, address.DB.Raw(query, email).Scan(&addressres).Error

}

func (address *AddressAdapter) AddAddress(addressreq helperstructs.AddressReq) (responce.AddressData, error) {

	var addressres responce.AddressData

	query := `INSERT INTO addresses (email,house_name,street_address,city,district,po,state) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id,email,house_name,street_address,city,district,po,state`

	return addressres, address.DB.Raw(query, addressreq.Email, addressreq.HouseName, addressreq.StreetAddress, addressreq.City, addressreq.District, addressreq.PO, addressreq.State).Scan(&addressres).Error

}

func (address *AddressAdapter) UpdateAddress(addressreq helperstructs.AddressReq) (responce.AddressData, error) {

	var addressres responce.AddressData

	query := `UPDATE addresses SET house_name = $1,street_address = $2,city = $3,district = $4,po = $5,state = $6 WHERE email = $7 AND id = $8 RETURNING id,email,house_name,street_address,city,district,po,state`

	return addressres, address.DB.Raw(query, addressreq.HouseName, addressreq.StreetAddress, addressreq.City, addressreq.District, addressreq.PO, addressreq.State, addressreq.Email, addressreq.ID).Scan(&addressres).Error

}

func (address *AddressAdapter) DeleteAddress(id uint) error {

	query := `DELETE FROM addresses where id = $1`

	return address.DB.Exec(query, id).Error

}
