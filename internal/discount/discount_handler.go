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
	discountBrandService *service.DiscountBrandService
	discountUserService  *service.DiscountUserService
}

func NewDiscountHandler(
	discountBrandService *service.DiscountBrandService,
	discountUserService *service.DiscountUserService,
) *DiscountHandler {
	return &DiscountHandler{
		discountBrandService: discountBrandService,
		discountUserService:  discountUserService,
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

	dh.discountBrandService.CreateDiscount(reqBrandDiscount)

	c.Writer.WriteHeader(200)
}

func (dh *DiscountHandler) CreateUserDiscountHandler(c *gin.Context) {
	var reqUserDiscount discount.RequestUserDiscount

	userId, err := strconv.Atoi(c.GetHeader("Blg-User-Id"))
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	reqUserDiscount.UserId = userId

	rntReq, err := dh.discountUserService.FetchByUserId(reqUserDiscount)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	c.IndentedJSON(http.StatusOK, rntReq)
}
