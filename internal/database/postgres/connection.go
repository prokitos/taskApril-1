package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	log "github.com/sirupsen/logrus"
)

func ConnectToDb(path string) *sql.DB {

	log.Info("connecting to the database")

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
	}

	db.Begin()

	return db
}

func MigrateStart() {

	duration := time.Second * 5
	time.Sleep(duration)

	db := ConnectToDb("internal/config/postgres.env")

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "internal/database/migrations"); err != nil {
		log.Error("migration connection error")
		panic(err)
	}
}
