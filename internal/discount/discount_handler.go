package discount

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hfleury/henrybillogram/pkg/discount"
	"github.com/hfleury/henrybillogram/pkg/discount/service"
)

type DiscountHandler struct {
	discountService *service.DiscountBrandService
}

func NewDiscountHandler(
	discountService *service.DiscountBrandService,
) *DiscountHandler {
	return &DiscountHandler{
		discountService: discountService,
	}
}

func (dh *DiscountHandler) CreateBrandDiscountHandler(c *gin.Context) {
	var reqBrandDiscount discount.RequestBrandDiscount

	if err := c.BindJSON(&reqBrandDiscount); err != nil {
		fmt.Printf("ERROR: %v", err)
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	brandId, err := strconv.Atoi(c.GetHeader("Blg-Brand-Id"))
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	reqBrandDiscount.BrandId = brandId
	reqBrandDiscount.BrandName = c.GetHeader("Blg-Brand-Name")

	dh.discountService.CreateDiscount(reqBrandDiscount)

	c.Writer.WriteHeader(200)
}
