package mysql

import (
	"database/sql"
	"log"
)

func ConnectionDB() *sql.DB {
	//　環境変数の設定
	db, err := sql.Open("mysql", "root:Shinya0023@tcp(127.0.0.1:3306)/fix_test?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
