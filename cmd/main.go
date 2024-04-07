package main

import (
	"module/internal/app"
	"module/internal/database"

	log "github.com/sirupsen/logrus"
)

// @title User API
// @version 1.0
// @description This is a sample service for managing users
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8888
// @BasePath /
func main() {

	log.SetLevel(log.DebugLevel) // показывать логи debug уровня
	log.Info("the server is starting")

	go database.MigrateStart()

	app.MainServer()
}
