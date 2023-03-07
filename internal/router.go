package internal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hfleury/henrybillogram/internal/app"
)

type BlgRouter struct {
	appBlgRouter *app.AppBillogram
}

func NewBlgRouter(
	appBlgRouter *app.AppBillogram,
) *BlgRouter {
	return &BlgRouter{
		appBlgRouter: appBlgRouter,
	}
}

func (hr *BlgRouter) ConfigRouter() *gin.Engine {
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	hr.appBlgRouter.GinEngine.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	return hr.appBlgRouter.GinEngine
}
