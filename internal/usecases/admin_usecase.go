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

func NewAdminUsecase(repo repositories.AdminRepo) usecasesinterface.AdminUsecaseInterface  {

	return &AdminUsecase{AdminRepo: repo}

}

func (admin *AdminUsecase) AdminLogin(ctx context.Context, adminreq helperstructs.AdminReq) (responce.AdminData, error) {
	
	return admin.AdminRepo.GetByEmail(adminreq)

}

func (admin *AdminUsecase) GetUsers(ctx context.Context) ([]responce.AdminsideUsersData,error) {

	return admin.AdminRepo.GetAllUsers()

}

func (admin *AdminUsecase) Reportuser(ctx context.Context, email string) (error) {

	return admin.AdminRepo.ReportUser(email)
	
}