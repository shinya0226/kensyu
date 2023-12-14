package mysql

import (
	"database/sql"
	"log"
)

func ConnectionDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Shinya0023@tcp(127.0.0.1:3306)/test_fix?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
