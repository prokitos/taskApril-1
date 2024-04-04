package app

import (
	"encoding/json"
	"io"
	"module/internal/models"
	"module/internal/services"
	"net/http"
)

func updateGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	reqBody, _ := io.ReadAll(r.Body)
	var curCar models.Car
	json.Unmarshal(reqBody, &curCar)

	services.RestUpdateData(w, &curCar)
}
