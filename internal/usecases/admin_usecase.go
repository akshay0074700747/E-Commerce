package usecases

import (
	"context"
	"ecommerce/internal/interfaces/repositories"
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type AdminUsecase struct {
	AdminRepo repositories.AdminRepo
}

func NewAdminUsecase(repo repositories.AdminRepo) usecasesinterface.AdminUsecaseInterface {

	return &AdminUsecase{AdminRepo: repo}

}

func (admin *AdminUsecase) AdminLogin(ctx context.Context, adminreq helperstructs.AdminReq) (responce.AdminData, error) {

	return admin.AdminRepo.GetByEmail(adminreq)

}

func (admin *AdminUsecase) GetUsers(ctx context.Context) ([]responce.AdminsideUsersData, error) {

	userdata, err := admin.AdminRepo.GetAllUsers()

	if err != nil {
		return userdata, err
	}

	for i := range userdata {

		count, err := admin.AdminRepo.GetReports(userdata[i].Email)

		if err != nil {
			return userdata, err
		}

		userdata[i].Reports = count
	}

	return userdata, nil

}

func (admin *AdminUsecase) Reportuser(ctx context.Context, reportreq helperstructs.ReportReq) error {

	return admin.AdminRepo.ReportUser(reportreq)

}

func (admin *AdminUsecase) GetUser(ctx context.Context, email string) (responce.AdminsideUsersData, error) {

	return admin.AdminRepo.GetUser(email)

}
