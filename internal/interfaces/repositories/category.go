package repositories

import (
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type CategoryRepo interface {
	GetallCategory()([]responce.CategoryData,error)
	CreateCategory(catreq helperstructs.CategoryReq) ( responce.CategoryData , error)
	UpdateCategory(catreq helperstructs.CategoryReq) ( responce.CategoryData , error)
	DeleteCategory(category_id uint) ( error)
}