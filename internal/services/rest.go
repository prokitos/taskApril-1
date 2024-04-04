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

	// здесь передача данных на сервер
	// сначала добавляем владельца
	// потом добавляем машину

}
