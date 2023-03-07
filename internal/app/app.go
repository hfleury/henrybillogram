package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hfleury/henrybillogram/internal/config"
)

type AppBillogram struct {
	EnvVars   map[string]string
	GinEngine *gin.Engine
}

func NewAppBillogram() (*AppBillogram, error) {
	app := &AppBillogram{}
	err := new(error)
	app.EnvVars, *err = config.SetEnv()
	app.GinEngine = gin.Default()

	return app, *err
}
