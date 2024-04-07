package postgres

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	log "github.com/sirupsen/logrus"
)

var envConnvertion string = "internal/config/postgres.env"
var migrationRoute string = "internal/database/migrations"

func ConnectToDb(w *http.ResponseWriter) *sql.DB {

	godotenv.Load(envConnvertion)

	envUser := os.Getenv("User")
	envPass := os.Getenv("Pass")
	envHost := os.Getenv("Host")
	envPort := os.Getenv("Port")
	envName := os.Getenv("Name")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", envUser, envPass, envHost, envPort, envName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error("database connection error")
		log.Debug("there is not connection with database")
		CheckError(err)
	}

	db.Begin()

	return db
}

func GetExternalRoutes(address *string, route *string) {
	godotenv.Load(envConnvertion)

	*address = os.Getenv("ExtAddress")
	*route = os.Getenv("ExtRoute")
}

func MigratonDbConnect(path string) *sql.DB {
	godotenv.Load(path)

	envUser := os.Getenv("User")
	envPass := os.Getenv("Pass")
	envHost := os.Getenv("Host")
	envPort := os.Getenv("Port")
	envName := os.Getenv("Name")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", envUser, envPass, envHost, envPort, envName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error("database connection error")
		log.Debug("there is not connection with database")
		CheckError(err)
	}

	db.Begin()

	return db
}

func MigrateStart() {

	duration := time.Second * 5
	time.Sleep(duration)

	db := MigratonDbConnect(envConnvertion)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Debug("dont get migration dialect")
		CheckError(err)
	}

	if err := goose.Up(db, migrationRoute); err != nil {
		log.Error("migration connection error")
		log.Debug("there is not connection with database in migration")
		CheckError(err)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
