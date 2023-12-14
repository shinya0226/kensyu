package mysql

import (
	"database/sql"
)

func ConnectionDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Shinya0023@tcp(127.0.0.1:3306)/yamamoto?parseTime=true")
	if err != nil {
		return db
	}
	return db
}
