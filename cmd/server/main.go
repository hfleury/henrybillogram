package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/hfleury/henrybillogram/internal"
	"github.com/hfleury/henrybillogram/internal/app"
	"github.com/hfleury/henrybillogram/internal/discount"
	"github.com/hfleury/henrybillogram/pkg/discount/repo"
	"github.com/hfleury/henrybillogram/pkg/discount/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// ctx := context.Background()
	appBlg, err := app.NewAppBillogram()
	if err != nil {
		log.Fatalln("Error initializing App")
	}

	// Graceful shutdown the server, wait remaings connections to finish
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - 15s")
	flag.Parse()

	// Init DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=CET",
		appBlg.EnvVars["BILLO_POSTGRESQL_HOST"],
		appBlg.EnvVars["BILLO_POSTGRESQL_USER"],
		appBlg.EnvVars["BILLO_POSTGRESQL_PASS"],
		appBlg.EnvVars["BILLO_POSTGRESQL_DB"],
		appBlg.EnvVars["BILLO_POSTGRESQL_PORT"],
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connectio to DB - error: %v", err)
	}

	// Init Repo
	dscRepo := repo.NewDiscountBrandRepo(db)
	dscUserRepo := repo.NewDiscountUserRepo(db)

	// Ini Services
	dscService := service.NewDiscountBrandService(*dscRepo)
	dscServiceUser := service.NewDiscountUserService(*dscUserRepo)

	// Init Handlers
	dscHandler := discount.NewDiscountHandler(dscService, dscServiceUser)

	// Init and config router
	blgRouter := internal.NewBlgRouter(appBlg)
	ginEngine := blgRouter.ConfigRouter()
	dscRoute := discount.NewDiscountRouter(appBlg, dscHandler)
	dscRoute.SetDiscountRouters()

	// Start the server
	hostUrl := fmt.Sprintf("%v:%v", appBlg.EnvVars["BILLO_HOST_ADDRESS"], appBlg.EnvVars["BILLO_HOST_PORT"])
	ginEngine.Run(hostUrl)
}
