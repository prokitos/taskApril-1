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

	// сначала надо получить айди владельца машины
	// потом удалить владельца машины
	// и уже потом удалить машину

	query := QueryCarDelete(id)
	connection := postgres.ConnectToDb(envConnvertion)
	postgres.ExecuteToDB(connection, w, query)
}

// Показать записи
func RestShowData(w http.ResponseWriter, curModel *models.Car, offset string, limit string, sort string) {

	// здесь передача данных на сервер
	var query string = "SELECT car.id,regnum,mark,model,year,people.id,name,surname,patronymic FROM car left join people on car.owner = people.id where car.id = 2"
	connection := postgres.ConnectToDb(envConnvertion)
	postgres.ShowFromDB(connection, w, query)

}

// Создать новую запись
func RestCreateData(w http.ResponseWriter, curModel *models.Car) {

	// здесь нужно обогащать данные
	var curOwner models.People
	curOwner.Name = "Andrew"
	curOwner.Surname = "Dyabov"
	curOwner.Patronymic = "Ivanovich"

	curModel.Mark = "Lada"
	curModel.Model = "Vesta"
	curModel.Year = "2002"
	curModel.RegNum = "x123xx150"
	curModel.Owner = curOwner

	var query string = "INSERT INTO people (name,surname,patronymic) VALUES ('thomas','Dyablow','ichvilnicht') RETURNING id"
	connection := postgres.ConnectToDb(envConnvertion)
	curOwner.Id = postgres.ExecuteReturnToDB(connection, w, query)

	query = "INSERT INTO car (regnum,mark,model,year,owner) VALUES ('xx100xx','lada','vesta','2000','6') RETURNING id"
	connection = postgres.ConnectToDb(envConnvertion)
	curOwner.Id = postgres.ExecuteReturnToDB(connection, w, query)

	// здесь передача данных на сервер
	// сначала добавляем владельца
	// потом добавляем машину

	// RETURNING id;
	// нельзя добавлять insertom в две таблицы!!!

}
