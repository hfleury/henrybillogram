package repo

import (
	"github.com/hfleury/henrybillogram/pkg/discount"
	"gorm.io/gorm"
)

type DiscountUserRepo struct {
	dbConn *gorm.DB
}

func NewDiscountUserRepo(
	dbConn *gorm.DB,
) *DiscountUserRepo {
	return &DiscountUserRepo{
		dbConn: dbConn,
	}
}

func (dbr *DiscountUserRepo) FetchDiscountByUserId(userId int) ([]discount.Discount, error) {
	var rtnDiscounts []discount.Discount
	dbr.dbConn.Where(discount.Discount{UserID: &userId}).Find(&rtnDiscounts)
	return rtnDiscounts, nil
}
