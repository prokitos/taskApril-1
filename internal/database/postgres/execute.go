package postgres

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// вызов операции над таблицей
func ExecuteToDB(db *sql.DB, w http.ResponseWriter, conn string) {
	defer db.Close()

	result, _ := db.Exec(conn)
	json.NewEncoder(w).Encode(result)

}

func ExecuteReturnToDB(db *sql.DB, w http.ResponseWriter, conn string) string {

	lastInsertId := "0"
	db.QueryRow(conn).Scan(&lastInsertId)
	return lastInsertId
}
