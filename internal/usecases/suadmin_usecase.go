package usecases

import (
	"context"
	"ecommerce/internal/interfaces/repositories"
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type SuAdminUsecase struct {
	SuAdminRepo repositories.SuAdminRepo
}

func NewSuAdminUsecase(repo repositories.SuAdminRepo) usecasesinterface.SuAdminUsecaseInterface {

	return &SuAdminUsecase{SuAdminRepo: repo}

}

func (supadmin *SuAdminUsecase) SuAdminLogin(ctx context.Context, suadmin helperstructs.SuAdminReq) (responce.SuAdminData, error) {

	return supadmin.SuAdminRepo.GetByEmail(suadmin)

}

func (supadmin *SuAdminUsecase) CreateAdmin(ctx context.Context, admin helperstructs.AdminReq) (responce.AdminData, error) {

	return supadmin.SuAdminRepo.CreateAdmin(admin)


}
