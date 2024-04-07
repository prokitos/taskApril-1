package services

import (
	"encoding/json"
	"io/ioutil"
	"module/internal/database/postgres"
	"module/internal/models"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

// Обновление записи
func RestUpdateData(w *http.ResponseWriter, curModel *models.Car) {

	connection := postgres.ConnectToDb(w)
	defer connection.Close()

	if curModel.Id == "" {
		return
	}

	query := ConStringUpdateCar(curModel)
	if query != "" {
		postgres.ExecuteToDB(connection, w, query)
	}

	query = ConStringUpdatePeople(curModel)
	if query != "" {
		postgres.ExecuteToDB(connection, w, query)
	}

}

// Удаление записи по айди
func RestDeleteData(w *http.ResponseWriter, id string) {

	connection := postgres.ConnectToDb(w)
	defer connection.Close()

	// получаем айдишник владельца машины
	var query string = "SELECT owner from car where id = " + id
	temp := postgres.ExecuteReturnToDB(connection, w, query)

	// каскадно удаляем владельца и машину
	query = "delete from people where id = " + temp
	postgres.ExecuteToDB(connection, w, query)

}

// Показать записи
func RestShowData(w *http.ResponseWriter, curModel *models.Car, offset string, limit string, sort string) {

	connection := postgres.ConnectToDb(w)
	defer connection.Close()

	// здесь передача данных на сервер
	var query string = ConStringShowSpec(curModel, limit, offset, sort)
	//"SELECT car.id,regnum,mark,model,year,people.id,name,surname,patronymic FROM car left join people on car.owner = people.id where car.id = 2"
	postgres.ShowFromDB(connection, w, query)

}

// Создать новую запись
func RestCreateData(w *http.ResponseWriter, regNums *models.CarNumber) {

	connection := postgres.ConnectToDb(w)
	defer connection.Close()

	for _, i := range regNums.RegNum {
		curCar := sendRequestToGet(i)

		var query string = ConStringInsertPeople(&curCar)
		curCar.Owner.Id = postgres.ExecuteReturnToDB(connection, w, query)

		query = ConStringInsertCar(&curCar)
		postgres.ExecuteToDB(connection, w, query)

	}

}

func sendRequestToGet(p_num string) models.Car {

	var strConnect string = ""
	var strRoute string = ""
	postgres.GetExternalRoutes(&strConnect, &strRoute)

	baseURL, _ := url.Parse(strConnect + "/info")
	params := url.Values{}
	params.Add("regNum", p_num)
	baseURL.RawQuery = params.Encode()

	resp, err := http.Get(baseURL.String())
	if err != nil {
		log.Debug("Error connecting to external api")
		log.Error("Error getting data from api")
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var carExemp models.Car
	json.Unmarshal(body, &carExemp)
	return carExemp
}
