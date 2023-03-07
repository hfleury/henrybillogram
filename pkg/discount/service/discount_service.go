package service

import (
	"math/rand"
	"strings"
	"sync"

	"github.com/hfleury/henrybillogram/pkg/discount"
	"github.com/hfleury/henrybillogram/pkg/discount/repo"
)

type DiscountBrandService struct {
	discountBrandRepo repo.DiscountBrandRepo
}

func NewDiscountBrandService(
	dscBrandRepo repo.DiscountBrandRepo,
) *DiscountBrandService {
	return &DiscountBrandService{
		discountBrandRepo: dscBrandRepo,
	}
}

func (dbs *DiscountBrandService) CreateDiscount(rbd discount.RequestBrandDiscount) {
	var discounts []discount.Discount
	wg := sync.WaitGroup{}

	for i := 0; i < rbd.Amount; i++ {
		wg.Add(1)
		go generateDiscountCode(&discounts, rbd.BrandName, rbd.BrandId, &wg)
	}
	wg.Wait()
	dbs.discountBrandRepo.InsertBatchDiscount(discounts)
}

func generateDiscountCode(discounts *[]discount.Discount, brandName string, brandId int, wg *sync.WaitGroup) {
	defer wg.Done()
	dsc := discount.Discount{
		BrandID: brandId,
		Code:    brandName + "_" + randomString(30),
	}
	*discounts = append(*discounts, dsc)
}

func randomString(n int) string {
	var alphabet []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

	alphabetSize := len(alphabet)
	var sb strings.Builder

	for i := 0; i < n; i++ {
		ch := alphabet[rand.Intn(alphabetSize)]
		sb.WriteRune(ch)
	}

	s := sb.String()
	return s
}
