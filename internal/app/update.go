package app

import (
	"encoding/json"
	"io"
	"module/internal/models"
	"module/internal/services"
	"net/http"
)

// Cars godoc
// @Summary Update cars and persons
// @Description Update cars and persons
// @Tags cars
// @Accept  json
// @Produce  json
// @Param user body models.Car true "Update cars"
// @Failure 400 "Invalid parameters supplied"
// @Failure 404 "cars or peoples not found"
// @Router /update [put]
func updateGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	reqBody, _ := io.ReadAll(r.Body)
	var curCar models.Car
	json.Unmarshal(reqBody, &curCar)

	services.RestUpdateData(&w, &curCar)
}
