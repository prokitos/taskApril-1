package main

import (
	"module/internal/app"
	"module/internal/database/postgres"

	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetLevel(log.DebugLevel) // показывать логи debug уровня
	log.Info("the server is starting")

	go postgres.MigrateStart()

	app.MainServer()
}
