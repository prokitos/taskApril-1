package main

import (
	"module/internal/app"

	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetLevel(log.DebugLevel) // показывать логи debug уровня
	log.Info("the server is starting")

	// здесь должна быть миграция

	app.MainServer()
}
