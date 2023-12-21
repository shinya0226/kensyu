package mysql

import (
	"database/sql"
)

func ConnectionDB() *sql.DB {
	//　環境変数の定義
	//　user := os.Getenv("DB_USER")
	//　pass := os.Getenv("DB_PASS")
	//　host := os.Getenv("DB_HOST")
	//　port := os.Getenv("DB_PORT")
	//　name := os.Getenv("DB_NAME")
	//　table := os.Getenv("DB_TABLE")
	//　DB読み込み
	//　dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)
	db, err := sql.Open("mysql", "root:Shinya0023@tcp(localhost:3306)/yamamoto?parseTime=true")
	if err != nil {
		return nil
	}
	return db
}
