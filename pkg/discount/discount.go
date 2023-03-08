package discount

import (
	"time"
)

type RequestBrandDiscount struct {
	Amount    int `json:"amountDiscount"`
	BrandName string
	BrandId   int
}

type RequestUserDiscount struct {
	UserId int
}

type Discount struct {
	ID        int        `gorm:"column:discount_id" json:"id"`
	BrandID   int        `gorm:"column:brand_id" json:"brandId"`
	UserID    *int       `gorm:"column:user_id" json:"userId"`
	Code      string     `gorm:"column:discount_code" json:"code"`
	CreatedAt time.Time  `gorm:"column:discount_created_at" json:"createdAt"`
	UsedAt    *time.Time `gorm:"column:discount_used_at" json:"usedAt"`
}

type DiscountBrandService interface {
	CreateDiscount(rbd RequestBrandDiscount)
}

type DiscountBrandRepo interface {
	InsertBatchDiscount(d []Discount) ([]Discount, error)
}

type DiscountUserRepo interface {
	FetchDiscountByUserId(userId int) ([]Discount, error)
}

func (Discount) TableName() string {
	return "discount"
}
