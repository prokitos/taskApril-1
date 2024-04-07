package postgres

import (
	"database/sql"
	"module/internal/models"
	"net/http"
)

// вызов операции над таблицей
func ExecuteToDB(db *sql.DB, w *http.ResponseWriter, conn string) {

	_, err := db.Exec(conn)
	CheckError(err)

	models.GoodResponse(w)

}

func ExecuteReturnToDB(db *sql.DB, w *http.ResponseWriter, conn string) string {

	lastInsertId := "0"
	db.QueryRow(conn).Scan(&lastInsertId)
	return lastInsertId
}
