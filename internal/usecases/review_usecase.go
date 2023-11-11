package usecases

import (
	"context"
	"ecommerce/internal/interfaces/repositories"
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
)

type ReviewUsecase struct {
	ReviewRepo repositories.ReviewRepo
}

func NewReviewUsecase(repo repositories.ReviewRepo) usecasesinterface.ReviewUsecaseInterface {
	return &ReviewUsecase{
		ReviewRepo: repo,
	}
}

func (review *ReviewUsecase) CreateReview(ctx context.Context, req helperstructs.ReviewReq) error {

	if err := review.ReviewRepo.VerifyOrderedUser(req.ReviewedBy, req.Product); err != nil {
		return err
	}

	return review.ReviewRepo.CreateReview(req)

}

func (review *ReviewUsecase) UpdateReview(ctx context.Context, req helperstructs.ReviewReq) error {

	return review.ReviewRepo.UpdateReview(req)

}

func (review *ReviewUsecase) GetReviewsByID(ctx context.Context, prodid uint) ([]responce.ReviewResponce, error) {

	return review.ReviewRepo.GetReviewsByID(prodid)

}

func (review *ReviewUsecase) GetReviwByEmail(ctx context.Context, email string) ([]responce.ReviewResponce, error) {

	return review.ReviewRepo.GetReviwByEmail(email)

}

func (review *ReviewUsecase) DeleteReview(ctx context.Context, id uint) error {

	return review.ReviewRepo.DeleteReview(id)

}
