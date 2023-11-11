package usecasesinterface

import (
	"context"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type AddessUsecaseInterface interface {
	GetallAddress(ctx context.Context,email string) ([]responce.AddressData, error)
	AddAddress(ctx context.Context,addressreq helperstructs.AddressReq) (responce.AddressData, error)
	UpdateAddress(ctx context.Context,addressreq helperstructs.AddressReq) (responce.AddressData, error)
	DeleteAddress(ctx context.Context,id uint) error
}