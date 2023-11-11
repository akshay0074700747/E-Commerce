package usecases

import (
	"context"
	"ecommerce/internal/interfaces/repositories"
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type AddressUsecase struct {
	AddressRepo repositories.AddressRepo
}

func NewAddressUsecase(repo repositories.AddressRepo) usecasesinterface.AddessUsecaseInterface {

	return &AddressUsecase{AddressRepo: repo}

}

func (address *AddressUsecase) GetallAddress(ctx context.Context, email string) ([]responce.AddressData, error) {

	return address.AddressRepo.GetallAddress(email)

}

func (address *AddressUsecase) AddAddress(ctx context.Context, addressreq helperstructs.AddressReq) (responce.AddressData, error) {

	return address.AddressRepo.AddAddress(addressreq)

}

func (address *AddressUsecase) UpdateAddress(ctx context.Context, addressreq helperstructs.AddressReq) (responce.AddressData, error) {

	return address.AddressRepo.UpdateAddress(addressreq)

}

func (address *AddressUsecase) DeleteAddress(ctx context.Context, id uint) error {

	return address.AddressRepo.DeleteAddress(id)

}
