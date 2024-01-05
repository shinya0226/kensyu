package mysql

import (
	"fmt"
	"os"

	"database/sql"
)

func ConnectionDB() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)
	return sql.Open("mysql", dsn)
}
