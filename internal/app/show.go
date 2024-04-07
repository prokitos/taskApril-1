package app

import (
	"module/internal/models"
	"module/internal/services"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// Cars godoc
// @Summary Show cars and person
// @Description Show cars and his owner
// @Tags cars
// @Accept  json
// @Produce  json
// @Param sort query string false "Sort records"
// @Param limit query int false "Show max limit records"
// @Param offset query int false "Show records with current offset"
// @Param user query carShow false "Update cars"
// @Failure 400 "Invalid parameters supplied"
// @Failure 404 "Cars not found"
// @Success 200 {object} carShow "successful operation"
// @Router /show [get]
func showsSpecGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var curCar models.Car
	curCar.Id = r.FormValue("id")
	curCar.RegNum = r.FormValue("regNum")
	curCar.Mark = r.FormValue("mark")
	curCar.Model = r.FormValue("model")
	curCar.Year = r.FormValue("year")
	curCar.Owner.Name = r.FormValue("name")
	curCar.Owner.Surname = r.FormValue("surname")
	curCar.Owner.Patronymic = r.FormValue("patronymic")

	var sorted string = r.FormValue("sort")
	var offset string = r.FormValue("offset")
	var limit string = r.FormValue("limit")

	// проверка что в offset,limit,id либо пустота либо числа.
	if _, err := strconv.Atoi(offset); offset != "" && err != nil {
		log.Error("the show request was not executed")
		log.Debug("offset params couldn't convert to a number, offset = " + offset)

		models.BadClientResponse(&w)
		return
	}
	if _, err := strconv.Atoi(limit); limit != "" && err != nil {
		log.Error("the show request was not executed")
		log.Debug("limit params couldn't convert to a number, limit = " + limit)

		models.BadClientResponse(&w)
		return
	}
	if _, err := strconv.Atoi(curCar.Id); curCar.Id != "" && err != nil {
		log.Error("the show request was not executed")
		log.Debug("id params couldn't convert to a number, id = " + curCar.Id)

		models.BadClientResponse(&w)
		return
	}

	services.RestShowData(&w, &curCar, offset, limit, sorted)
}

type carShow struct {
	Id     string     `json:"id" example:"12"`
	RegNum string     `json:"regNum" example:"x16xx150"`
	Mark   string     `json:"mark" example:"lada"`
	Model  string     `json:"model" example:"kalina"`
	Year   string     `json:"year" example:"2000"`
	Owner  peopleShow `json:"owner"`
}

type peopleShow struct {
	Id         string `json:"id" example:"14"`
	Name       string `json:"name" example:"james"`
	Surname    string `json:"surname" example:"johnson"`
	Patronymic string `json:"patronymic" example:"petrovich"`
}
