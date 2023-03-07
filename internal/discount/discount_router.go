package discount

import (
	"github.com/hfleury/henrybillogram/internal/app"
	"github.com/hfleury/henrybillogram/pkg/blgauth"
)

type DiscountRouter struct {
	appBlgRouter    *app.AppBillogram
	discountHandler *DiscountHandler
}

func NewDiscountRouter(
	appBlg *app.AppBillogram,
	discountHandler *DiscountHandler,
) *DiscountRouter {
	return &DiscountRouter{
		appBlgRouter:    appBlg,
		discountHandler: discountHandler,
	}
}

func (dr *DiscountRouter) SetDiscountRouters() {
	branGroup := dr.appBlgRouter.GinEngine.Group("/brand")

	branGroup.Use(blgauth.BrandAuthRequired())
	{
		branGroup.POST("/discount", dr.discountHandler.CreateBrandDiscountHandler)
	}
}
