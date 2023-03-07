package repo

import (
	"github.com/hfleury/henrybillogram/pkg/discount"
	"gorm.io/gorm"
)

type DiscountBrandRepo struct {
	dbConn *gorm.DB
}

func NewDiscountBrandRepo(
	dbConn *gorm.DB,
) *DiscountBrandRepo {
	return &DiscountBrandRepo{
		dbConn: dbConn,
	}
}

func (dbr *DiscountBrandRepo) InsertBatchDiscount(d []discount.Discount) ([]discount.Discount, error) {
	discounteCreated := dbr.dbConn.Create(&d)
	if discounteCreated.Error != nil {
		return []discount.Discount{}, discounteCreated.Error
	}

	return d, nil
}
