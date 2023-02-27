package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hfleury/henrybillogram/internal/config"
)

type AppDiscount struct {
	EnvVars   map[string]string
	GinEngine *gin.Engine
}

func NewApp() (*AppDiscount, error) {
	app := &AppDiscount{}
	err := new(error)
	app.EnvVars, *err = config.SetEnv()
	app.GinEngine = gin.Default()

	return app, *err
}
