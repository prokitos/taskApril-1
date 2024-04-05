package postgres

import (
	"database/sql"
	"encoding/json"
	"module/internal/models"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// показать таблицу
func ShowFromDB(db *sql.DB, w http.ResponseWriter, stroka string) {

	defer db.Close()

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

	json.NewEncoder(w).Encode(cars)

}
