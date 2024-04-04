package postgres

import (
	"database/sql"
	"net/http"
)

// вызов операции над таблицей
func ExecuteToDB(db *sql.DB, w http.ResponseWriter, conn string) {
	defer db.Close()

}
