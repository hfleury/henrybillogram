package internal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hfleury/henrybillogram/internal/app"
)

type DiscounteRouter struct {
	appDiscount *app.AppDiscount
}

func NewDiscounteRouter(
	appDiscount *app.AppDiscount,
) *DiscounteRouter {
	return &DiscounteRouter{
		appDiscount: appDiscount,
	}
}

func (hr *DiscounteRouter) ConfigRouter() *gin.Engine {
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	hr.appDiscount.GinEngine.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	return hr.appDiscount.GinEngine
}
