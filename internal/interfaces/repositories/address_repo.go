package repositories

import (
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type AddressRepo interface {
	GetallAddress(email string) ([]responce.AddressData, error)
	AddAddress(addressreq helperstructs.AddressReq) (responce.AddressData, error)
	UpdateAddress(addressreq helperstructs.AddressReq) (responce.AddressData, error)
	DeleteAddress(id uint) error
}
