package mysql

import (
	"database/sql"
	"log"
)

// DBに接続
func ConnectionDB() *sql.DB {
	//　環境変数の設定
	// dbUser := os.Getenv("DB_USER")
	// dbPass := os.Getenv("DB_PASS")
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_NAME")

	//　DB読み込み
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)
	// db, err := sql.Open("mysql", dsn)
	db, err := sql.Open("mysql", "root:Shinya0023@tcp(127.0.0.1:3306)/yamamoto?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}
	return db
}
