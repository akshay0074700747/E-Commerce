package usecasesinterface

import (
	"context"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type ReviewUsecaseInterface interface {
	CreateReview(context.Context, helperstructs.ReviewReq) error
	UpdateReview(context.Context, helperstructs.ReviewReq) error
	GetReviewsByID(ctx context.Context, prodid uint) ([]responce.ReviewResponce, error)
	GetReviwByEmail(ctx context.Context, email string) ([]responce.ReviewResponce, error)
	DeleteReview(ctx context.Context, id uint) error
}
