package mysql

import (
	"database/sql"
)

func ConnectionDB() *sql.DB {
	db, _ := sql.Open("mysql", "atsuser:atspass@tcp(localhost:3306)/kensyu_testing?parseTime=true")
	return db
}
