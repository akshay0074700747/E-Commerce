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

func (supadmin *SuAdminUsecase) BlockUser(ctx context.Context, blockreq helperstructs.BlockReq) error {

	return supadmin.SuAdminRepo.BlockUser(blockreq)

}

func (supadmin *SuAdminUsecase) GetAllUsers(ctx context.Context) ([]responce.AdminsideUsersData, error) {

	userdata, err := supadmin.SuAdminRepo.GetAllUsers()

	if err != nil {
		return userdata, err
	}

	for i := range userdata {

		count, err := supadmin.SuAdminRepo.GetReports(userdata[i].Email)

		if err != nil {
			return userdata, err
		}

		userdata[i].Reports = count
	}

	return userdata, nil

}

func (supadmin *SuAdminUsecase) GetAllAdmins(ctx context.Context) ([]responce.AdminData, error) {

	admins, err := supadmin.SuAdminRepo.GetAllAdmins()

	if err != nil {
		return admins, err
	}

	for i := range admins {

		count, err := supadmin.SuAdminRepo.GetReports(admins[i].Email)

		if err != nil {
			return admins, err
		}

		admins[i].Reports = count

	}

	return admins, nil

}

func (supadmin *SuAdminUsecase) GetReportes(ctx context.Context) ([]responce.DetailReportResponce, error) {

	return supadmin.SuAdminRepo.GetReportes()

}

func (supadmin *SuAdminUsecase) GetDetailedReport(ctx context.Context, email string) (responce.DetailReportResponce, error) {

	return supadmin.SuAdminRepo.GetDetailedReport(email)

}
