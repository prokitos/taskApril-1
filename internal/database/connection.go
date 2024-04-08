package database

import (
	"database/sql"
	"fmt"
	"module/internal/models"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	log "github.com/sirupsen/logrus"
)

// путь до .env миграций
var envConnvertion string = "internal/config/postgres.env"
var migrationRoute string = "internal/database/migrations"

// хранение соединения с базой данных
var curDbRef *sql.DB

// функция для получения базы данных из переменной
func GetConnection() *sql.DB {
	return curDbRef
}

// установка соединения с базой данных
func ConnectToDb() {

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
		models.CheckError(err)
	}

	db.Begin()

	curDbRef = db
}

func CloseConnectToDb() {
	curDbRef.Close()
}

// получить адрес внешнего сервера
func GetExternalRoutes(address *string) {
	godotenv.Load(envConnvertion)
	*address = os.Getenv("ExtAddress")
}

// начать миграцию
func MigrateStart() {

	duration := time.Second * 5
	time.Sleep(duration)

	db := GetConnection()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Debug("dont get migration dialect")
		models.CheckError(err)
	}

	if err := goose.Up(db, migrationRoute); err != nil {
		log.Error("migration connection error")
		log.Debug("no connection with database, or wrong migration route")
		models.CheckError(err)
	}
}
