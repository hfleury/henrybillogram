package config

import (
	"os"

	"github.com/joho/godotenv"
)

func SetEnv() (map[string]string, error) {
	runEnv := os.Getenv("BILLO_ENVIRONMENT")
	if runEnv != "" && runEnv == "LOCAL" {
		err := godotenv.Load(".env")
		if err != nil {
			return nil, err
		}
	}

	m := GetEnv()
	return m, nil
}

func GetEnv() map[string]string {
	m := make(map[string]string)
	// HOST env vars
	m["BILLO_HOST_PORT"] = os.Getenv("BILLO_HOST_PORT")
	m["BILLO_HOST_ADDRESS"] = os.Getenv("BILLO_HOST_ADDRESS")

	// DB env vars
	m["BILLO_POSTGRESQL_PORT"] = os.Getenv("BILLO_POSTGRESQL_PORT")
	m["BILLO_POSTGRESQL_DB"] = os.Getenv("BILLO_POSTGRESQL_DB")
	m["BILLO_POSTGRESQL_PASS"] = os.Getenv("BILLO_POSTGRESQL_PASS")
	m["BILLO_POSTGRESQL_HOST"] = os.Getenv("BILLO_POSTGRESQL_HOST")
	m["BILLO_POSTGRESQL_USER"] = os.Getenv("BILLO_POSTGRESQL_USER")

	return m
}
