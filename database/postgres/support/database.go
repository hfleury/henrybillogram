package support

import (
	"os"
	"testing"

	_ "github.com/golang-migrate/migrate/source/file"
	log "github.com/sirupsen/logrus"
	psql "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbConn *gorm.DB

// type dockerDBConn struct {
// 	GormConn *gorm.DB
// }
//
// var (
// 	// DockerDBConn holds the connection to our DB in the container we spin up for testing.
// 	DockerDBConn *dockerDBConn
// )

func TestMain(m *testing.M) {
	dbConn()
	// pool, resource := initDB()
	//initDB()
	code := m.Run()
	// closeDB(pool, resource)
	os.Exit(code)
}

func dbConn() {
	var err error

	dsn := "host=localhost user=rootuser password=nosecret dbname=billodb port=5432 sslmode=disable TimeZone=CET"
	DbConn, err = gorm.Open(psql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connectio to DB - error: %v", err)
	}
}

// func initMigrations(dbConn *gorm.DB) {
// 	sqlDB, err := dbConn.DB()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	migrate, err := migrate.NewWithDatabaseInstance(
// 		"file://../../../database/postgres/migration",
// 		"horsedbtest", driver)
// 	if err != nil {
// 		log.Fatalf("NewWithDatabaseInstance %v", err)
// 	}
//
// 	err = migrate.Up()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func initDB() (*dockertest.Pool, *dockertest.Resource) {
// 	pgURL := initPostgres()
// 	pgPass, _ := pgURL.User.Password()
//
// 	runOpts := dockertest.RunOptions{
// 		Repository: "postgres",
// 		Tag:        "15.1-alpine",
// 		Env: []string{
// 			"POSTGRES_USER=" + pgURL.User.Username(),
// 			"POSTGRES_PASSWORD=" + pgPass,
// 			"POSTGRES_DB=" + pgURL.Path,
// 		},
// 		ExposedPorts: []string{"5432"},
// 		PortBindings: map[docker.Port][]docker.PortBinding{
// 			"5432": {
// 				{HostIP: "localhost", HostPort: "5433"},
// 			},
// 		},
// 	}
//
// 	pool, err := dockertest.NewPool("")
// 	if err != nil {
// 		log.WithError(err).Fatal("Could not connect to docker")
// 	}
//
// 	resource, err := pool.RunWithOptions(&runOpts)
// 	if err != nil {
// 		log.WithError(err).Fatal("Could not start postgres container")
// 	}
//
// 	pgURL.Host = resource.Container.NetworkSettings.IPAddress
//
// 	// Docker layer network is different on Mac
// 	if runtime.GOOS == "darwin" {
// 		pgURL.Host = net.JoinHostPort(resource.GetBoundIP("5432/tcp"), resource.GetPort("5432/tcp"))
// 	}
//
// 	DockerDBConn = &dockerDBConn{}
// 	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
// 	if err := pool.Retry(func() error {
// 		DockerDBConn.GormConn, err = gorm.Open(psql.Open(pgURL.String()), &gorm.Config{})
// 		if err != nil {
// 			return err
// 		}
// 		return err
// 	}); err != nil {
// 		phrase := fmt.Sprintf("Could not connect to docker: %s", err)
// 		log.Error(phrase)
// 	}
//
// 	DockerDBConn.initMigrations()
//
// 	return pool, resource
// }

// func closeDB(pool *dockertest.Pool, resource *dockertest.Resource) {
// 	if err := pool.Purge(resource); err != nil {
// 		phrase := fmt.Sprintf("Could not purge resource: %s", err)
// 		log.Error(phrase)
// 	}
// }

// func (db dockerDBConn) initMigrations() {
// 	sqlDB, err := db.GormConn.DB()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	migrate, err := migrate.NewWithDatabaseInstance(
// 		"file://../../../database/postgres/migration",
// 		"horsedbtest", driver)
// 	if err != nil {
// 		log.Fatalf("NewWithDatabaseInstance %v", err)
// 	}
//
// 	err = migrate.Up()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func initPostgres() *url.URL {
// 	pgURL := &url.URL{
// 		Scheme: "postgres",
// 		User:   url.UserPassword("rootuser", "nosecret"),
// 		Path:   "horsedbtest",
// 		Host:   "localhost:5432",
// 	}
// 	q := pgURL.Query()
// 	q.Add("sslmode", "disable")
// 	pgURL.RawQuery = q.Encode()
//
// 	return pgURL
// }
