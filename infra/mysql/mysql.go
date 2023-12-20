package mysql

import (
	"database/sql"
	"fmt"
)

func ConnectionDB() *sql.DB {
	//　環境変数の定義
	//user := os.Getenv("DB_USER")
	//pass := os.Getenv("DB_PASS")
	//host := os.Getenv("DB_HOST")
	//port := os.Getenv("DB_PORT")
	//name := os.Getenv("DB_NAME")
	//　DB読み込み
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)
	dsn := fmt.Sprintf("atsuser:atspass@tcp(localhost:3306)/kensyu_testing?parseTime=true")

	db, _ := sql.Open("mysql", dsn)
	return db
}
