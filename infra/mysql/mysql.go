package mysql

import (
	"database/sql"
)

func ConnectionDB() *sql.DB {
	//　dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)
	db, err := sql.Open("mysql", "root:Shinya0023@tcp(localhost:3306)/test_fix?parseTime=true")
	if err != nil {
		return nil
	}
	return db
}
