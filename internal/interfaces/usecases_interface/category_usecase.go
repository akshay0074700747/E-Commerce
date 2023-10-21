package usecasesinterface

import (
	"context"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type CategoryUsecaseInterface interface {
	CreateCategory(ctx context.Context, catreq helperstructs.CategoryReq) (responce.CategoryData, error)
	UpdateCategory(ctx context.Context, catreq helperstructs.CategoryReq) (responce.CategoryData, error)
	DeleteCategory(ctx context.Context, cat_id uint) error
	GetAllCategories(ctx context.Context) ([]responce.CategoryData, error)
}
