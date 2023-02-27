package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"testing"
	"time"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	dc "github.com/ory/dockertest/v3/docker"
)

var (
	dbUser     = "rootuser"
	dbPassword = "nosecret"
	dbName     = "billodb"
	dbPort     = "5432"
	dbDsn      = "postgres://%s:%s@localhost:%s/%s?sslmode=disable"
)

func TestMain(m *testing.M) {
	migrationPath, err := filepath.Abs("../../database/postgres/migration")
	if err != nil {
		log.Fatalln("filepath.Abs() err:", err)
	}

	resource, pool := dockerTestDb()

	if err := runMigrations(migrationPath); err != nil {
		log.Fatalln("Error running runMigrations err:", err)
	}

	defer pool.Purge(resource)
}

func dockerTestDb() (*dockertest.Resource, *dockertest.Pool) {
	// Connection to the local Docker API enabling the creating,
	// running or deleting of Docker images.
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("could not connect to docker - %s", err.Error())
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	dcoptions := &dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "15.1-alpine",
		Env: []string{
			"POSTGRES_USER=" + dbUser,
			"POSTGRES_PASSWORD=" + dbPassword,
			"POSTGRES_DB=" + dbName,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[dc.Port][]dc.PortBinding{
			"5432": {
				{HostIP: "localhost", HostPort: dbPort},
			},
		},
	}

	resource, err := pool.RunWithOptions(dcoptions)

	if err != nil {
		log.Fatalf("could not build and run the docker - %s", err.Error())
	}

	return resource, pool
}

func runMigrations(migrationsPath string) error {
	if migrationsPath == "" {
		return errors.New("missing migrations path")
	}

	dsn := fmt.Sprintf(dbDsn, dbUser, dbPassword, dbPort, dbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln("error opening database driver:", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		time.Sleep(2 * time.Second)
		driver, err = postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			log.Fatalln("error postgres WithInstance:", err)
		}
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsPath,
		dbName, driver)
	if err != nil {
		log.Fatalln("error creating a new database instance:", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	m.Close()
	return nil
}

// func ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
// 	rr := httptest.NewRecorder()
// 	a, err := app.NewApp()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	return rr
// }
