package service

import (
	"errors"

	"github.com/hfleury/henrybillogram/pkg/discount"
	"github.com/hfleury/henrybillogram/pkg/discount/repo"
)

type DiscountUserService struct {
	discountUserRepo repo.DiscountUserRepo
}

func NewDiscountUserService(
	discountUserRepo repo.DiscountUserRepo,
) *DiscountUserService {
	return &DiscountUserService{
		discountUserRepo: discountUserRepo,
	}
}

func (dus *DiscountUserService) FetchByUserId(rud discount.RequestUserDiscount) ([]discount.Discount, error) {
	if rud.UserId <= 0 {
		return []discount.Discount{}, errors.New("User Id is required")
	}

	discounts, err := dus.discountUserRepo.FetchDiscountByUserId(rud.UserId)
	if err != nil {
		return []discount.Discount{}, nil
	}

	return discounts, err
}
