package repo

import (
	"log"
	"testing"

	"github.com/hfleury/henrybillogram/pkg/discount"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestDiscountUserRepo_FetchDiscountByUserId(t *testing.T) {

	dsn := "host=localhost user=rootuser password=nosecret dbname=billodb port=5432 sslmode=disable"
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connect to db: %v", err)
	}

	t.Run("Success get user discounts", func(t *testing.T) {
		discount := discount.Discount{
			BrandID: 1,
			UserID:  func(val int) *int { return &val }(1),
			Code:    "brand1_usertest1",
		}

		dbConn.Create(&discount)

		userRepo := NewDiscountUserRepo(dbConn)

		rtnTest, err := userRepo.FetchDiscountByUserId(1)
		assert.NoError(t, err)
		assert.Equal(t, discount.BrandID, rtnTest[0].BrandID)
		assert.Equal(t, discount.UserID, rtnTest[0].UserID)
	})
}
