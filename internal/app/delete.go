package app

import (
	"module/internal/models"
	"module/internal/services"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// Cars godoc
// @Summary Delete car and person by car id
// @Description Delete car and perso by car id
// @Tags cars
// @Accept  json
// @Produce  json
// @Param id query int true "delete car by id"
// @Failure 400 "Invalid id supplied"
// @Failure 404 "Car not found"
// @Router /delete [delete]
func deleteRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// получение айди, и если айди не может конвертится в число, то выдаем ошибку
	id := r.FormValue("id")
	if _, err := strconv.Atoi(id); err != nil {
		log.Debug("id couldn't convert to a number: " + id)
		models.BadClientResponse(&w)
		return
	}

	services.RestDeleteData(&w, id)
}
