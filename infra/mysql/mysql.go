package mysql

import (
	"database/sql"
	"log"
)

func ConnectionDB() *sql.DB {
	db, err := sql.Open("mysql", "atsuser:atspass@tcp(localhost:3306)/kensyu_testing?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
