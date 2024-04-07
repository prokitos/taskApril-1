package services

import (
	"encoding/json"
	"io"
	"module/internal/database/dao"
	"module/internal/models"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// обновление записи
func CarUpdate(w *http.ResponseWriter, r *http.Request) {

	reqBody, _ := io.ReadAll(r.Body)
	var curCar models.Car
	json.Unmarshal(reqBody, &curCar)

	dao.UpdateData(w, &curCar)
}

// Удаление записи по айди
func CarDelete(w *http.ResponseWriter, r *http.Request) {

	// получение айди, и если айди не может конвертится в число, то выдаем ошибку
	id := r.FormValue("id")
	if _, err := strconv.Atoi(id); err != nil {
		log.Debug("id couldn't convert to a number: " + id)
		models.BadClientResponse(w)
		return
	}

	dao.DeleteData(w, id)
}

// Показать записи
func CarShow(w *http.ResponseWriter, r *http.Request) {

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

		models.BadClientResponse(w)
		return
	}
	if _, err := strconv.Atoi(limit); limit != "" && err != nil {
		log.Error("the show request was not executed")
		log.Debug("limit params couldn't convert to a number, limit = " + limit)

		models.BadClientResponse(w)
		return
	}
	if _, err := strconv.Atoi(curCar.Id); curCar.Id != "" && err != nil {
		log.Error("the show request was not executed")
		log.Debug("id params couldn't convert to a number, id = " + curCar.Id)

		models.BadClientResponse(w)
		return
	}

	dao.ShowData(w, &curCar, offset, limit, sorted)
}

// Создать новую запись
func CarCreate(w *http.ResponseWriter, r *http.Request) {

	reqBody, _ := io.ReadAll(r.Body)
	var nomera models.CarNumber
	json.Unmarshal(reqBody, &nomera)

	for _, i := range nomera.RegNum {
		curCar := sendRequestToGet(i)
		dao.CreateData(w, &curCar)
	}

}
