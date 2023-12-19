package mysql

import (
	"database/sql"
	"fmt"
	"os"
)

func ConnectionDB() *sql.DB {
	//　環境変数の定義
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")
	//　DB読み込み
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DBUser, DBPass, DBHost, DBPort, DBName)
	db, _ := sql.Open("mysql", dsn)
	// db, _ := sql.Open("mysql", "atsuser:atspass@tcp(localhost:3306)/kensyu_testing?parseTime=true")
	return db
}
