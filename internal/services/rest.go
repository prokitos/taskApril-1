package services

import (
	"module/internal/database/postgres"
	"module/internal/models"
	"net/http"
)

var envConnvertion string = "internal/config/postgres.env"

// Обновление записи
func RestUpdateData(w http.ResponseWriter, curModel *models.Car) {

	// здесь передача данных на сервер

}

// Удаление записи по айди
func RestDeleteData(w http.ResponseWriter, id string) {

	connection := postgres.ConnectToDb(envConnvertion)
	defer connection.Close()

	// получаем айдишник владельца машины
	var query string = "SELECT owner from car where id = " + id
	temp := postgres.ExecuteReturnToDB(connection, w, query)

	// каскадно удаляем владельца и машину
	query = "delete from people where id = " + temp
	postgres.ExecuteToDB(connection, w, query)

}

// Показать записи
func RestShowData(w http.ResponseWriter, curModel *models.Car, offset string, limit string, sort string) {

	connection := postgres.ConnectToDb(envConnvertion)
	defer connection.Close()

	// здесь передача данных на сервер
	var query string = ConStringShowSpec(curModel, limit, offset)
	//"SELECT car.id,regnum,mark,model,year,people.id,name,surname,patronymic FROM car left join people on car.owner = people.id where car.id = 2"
	postgres.ShowFromDB(connection, w, query)

}

// Создать новую запись
func RestCreateData(w http.ResponseWriter, regNums *models.CarNumber) {

	connection := postgres.ConnectToDb(envConnvertion)
	defer connection.Close()

	// здесь нужно обогащать данные
	var curOwner models.People
	curOwner.Name = "Nikita"
	curOwner.Surname = "Dyabov"
	curOwner.Patronymic = "Ivanovich"

	var curCar models.Car
	curCar.Mark = "Lada"
	curCar.Model = "Verdana"
	curCar.Year = "2000"
	curCar.RegNum = "x123xx150"
	curCar.Owner = curOwner

	// var query string = "INSERT INTO people (name,surname,patronymic) VALUES ('Andrew','Dyablow','ichvilnicht') RETURNING id"
	// connection := postgres.ConnectToDb(envConnvertion)
	// curOwner.Id = postgres.ExecuteReturnToDB(connection, w, query)

	// query = "INSERT INTO car (regnum,mark,model,year,owner) VALUES ('xx100xx','mercedes','s666','2010','11') RETURNING id"
	// connection = postgres.ConnectToDb(envConnvertion)
	// curOwner.Id = postgres.ExecuteReturnToDB(connection, w, query)

}
