package mysql

import (
	"database/sql"
)

func ConnectionDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "atsuser:atspass@tcp(localhost:3306)/kensyu_testing?parseTime=true")
	if err != nil {
		return db, err
	}
	return db, nil
}
