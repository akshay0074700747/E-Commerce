package usecasesinterface

import (
	"context"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type ProductUsecaseInterface interface {
	AddProduct(ctx context.Context, productreq helperstructs.ProductReq) (responce.ProuctData,error)
	GetProducts(ctx context.Context)([]responce.ProuctData,error)
	UpdateProduct(ctx context.Context, productreq helperstructs.ProductReq) (responce.ProuctData,error)
	DeleteProduct(ctx context.Context, product_id uint ) (error)
}