package postgres

import (
	"database/sql"
	"net/http"
)

// показать таблицу
func ShowFromDB(db *sql.DB, w http.ResponseWriter, stroka string) {

	defer db.Close()

}
