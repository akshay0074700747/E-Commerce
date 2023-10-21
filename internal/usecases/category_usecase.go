package usecases

import (
	"context"
	"ecommerce/internal/interfaces/repositories"
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type CategoryUsecase struct {
	CategoryRepo repositories.CategoryRepo
}

func NewCategoryUsecase(repo repositories.CategoryRepo) usecasesinterface.CategoryUsecaseInterface {

	return &CategoryUsecase{CategoryRepo: repo}

}

func (cat *CategoryUsecase) CreateCategory(ctx context.Context, catreq helperstructs.CategoryReq) (responce.CategoryData, error) {

	return cat.CategoryRepo.CreateCategory(catreq)

}

func (cat *CategoryUsecase) UpdateCategory(ctx context.Context, catreq helperstructs.CategoryReq) (responce.CategoryData, error) {

	return cat.CategoryRepo.UpdateCategory(catreq)

}

func (cat *CategoryUsecase) DeleteCategory(ctx context.Context, cat_id uint) error {

	return cat.CategoryRepo.DeleteCategory(cat_id)

}

func (cat *CategoryUsecase) GetAllCategories(ctx context.Context) ([]responce.CategoryData, error) {

	return cat.CategoryRepo.GetallCategory()

}
