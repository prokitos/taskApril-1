package app

import (
	"module/internal/models"
	"module/internal/services"
	"net/http"
)

func showsSpecGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var curCar models.Car
	curCar.Id = r.FormValue("id")
	curCar.RegNum = r.FormValue("regNum")
	curCar.Mark = r.FormValue("mark")
	curCar.Model = r.FormValue("model")
	curCar.Year = r.FormValue("year")

	var sorted string = r.FormValue("sort")
	var offset string = r.FormValue("offset")
	var limit string = r.FormValue("limit")

	services.RestShowData(w, &curCar, offset, limit, sorted)
}
