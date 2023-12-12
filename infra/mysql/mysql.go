package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// DBに接続
func ConnectionDB() *sql.DB {
	//　環境変数の設定
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")

	//　DB読み込み
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", db_user, db_pass, db_host, db_port, db_name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
