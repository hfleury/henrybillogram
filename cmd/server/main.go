package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/hfleury/henrybillogram/internal"
	"github.com/hfleury/henrybillogram/internal/app"
)

func main() {
	// ctx := context.Background()
	app, err := app.NewApp()
	if err != nil {
		log.Fatalln("Error initializing App")
	}
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// Config router
	discountRouter := internal.NewDiscounteRouter(app)
	ginEngine := discountRouter.ConfigRouter()

	// TODO: make a pkg or put it in a specific file
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=CET",
	// 	app.EnvVars["HORSE_POSTGRESQL_HOST"],
	// 	app.EnvVars["HORSE_POSTGRESQL_USER"],
	// 	app.EnvVars["HORSE_POSTGRESQL_PASS"],
	// 	app.EnvVars["HORSE_POSTGRESQL_DB"],
	// 	app.EnvVars["HORSE_POSTGRESQL_PORT"],
	// )
	// _, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }

	// Init Repo

	// Start the server
	hostUrl := fmt.Sprintf("%v:%v", app.EnvVars["HORSE_HOST_ADDRESS"], app.EnvVars["HORSE_HOST_PORT"])
	ginEngine.Run(hostUrl)
}
