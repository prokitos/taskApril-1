package dao

import (
	"database/sql"
	"encoding/json"
	"module/internal/database"
	"module/internal/models"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Обновление записи
func UpdateData(w *http.ResponseWriter, curModel *models.Car) {

	connection := database.ConnectToDb(w)
	defer connection.Close()

	if curModel.Id == "" {
		return
	}

	query := conStringUpdateCar(curModel)
	if query != "" {
		executeToDB(connection, w, query)
	}

	query = conStringUpdatePeople(curModel)
	if query != "" {
		executeToDB(connection, w, query)
	}

}

// Удаление записи по айди
func DeleteData(w *http.ResponseWriter, id string) {

	connection := database.ConnectToDb(w)
	defer connection.Close()

	// получаем айдишник владельца машины
	var query string = "SELECT owner from car where id = " + id
	temp := executeReturnToDB(connection, w, query)

	// каскадно удаляем владельца и машину
	query = "delete from people where id = " + temp
	executeToDB(connection, w, query)

}

// Показать записи
func ShowData(w *http.ResponseWriter, curModel *models.Car, offset string, limit string, sort string) {

	connection := database.ConnectToDb(w)
	defer connection.Close()

	// здесь передача данных на сервер
	var query string = conStringShowSpec(curModel, limit, offset, sort)
	//"SELECT car.id,regnum,mark,model,year,people.id,name,surname,patronymic FROM car left join people on car.owner = people.id where car.id = 2"
	showFromDB(connection, w, query)

}

// Создать новую запись
func CreateData(w *http.ResponseWriter, curCar *models.Car) {

	connection := database.ConnectToDb(w)
	defer connection.Close()

	var query string = conStringInsertPeople(curCar)
	curCar.Owner.Id = executeReturnToDB(connection, w, query)

	query = conStringInsertCar(curCar)
	executeToDB(connection, w, query)

}

// вызов операции над таблицей
func executeToDB(db *sql.DB, w *http.ResponseWriter, conn string) {

	_, err := db.Exec(conn)
	CheckError(err)

	models.GoodResponse(w)

}

func executeReturnToDB(db *sql.DB, w *http.ResponseWriter, conn string) string {

	lastInsertId := "0"
	db.QueryRow(conn).Scan(&lastInsertId)
	return lastInsertId
}

// показать таблицу
func showFromDB(db *sql.DB, w *http.ResponseWriter, stroka string) {

	rows, _ := db.Query(stroka)
	defer rows.Close()

	cars := []models.Car{}

	for rows.Next() {
		car := models.Car{}
		own := models.People{}
		err := rows.Scan(&car.Id, &car.RegNum, &car.Model, &car.Mark, &car.Year, &own.Id, &own.Name, &own.Surname, &own.Patronymic)
		if err != nil {
			log.Error("database show data error")
			log.Debug("errored query when show data: " + stroka)

			continue
		}
		car.Owner = own
		cars = append(cars, car)
	}

	json.NewEncoder(*w).Encode(cars)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
