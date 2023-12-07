package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

// DBに接続
func ConnectionDB() *sql.DB {
	// //環境変数の設定
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")

	//DB読み込み
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DBUser, DBPass, DBHost, DBPort, DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(http.StatusInternalServerError)
		log.Fatal(err)
	}
	return db
}
