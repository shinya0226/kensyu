package main

import (
	"fmt"
	"os"
	
	"log"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/testfixtures.v1"
)

func main() {
	prepareTestDatabase()
}

const FixturesPath = "testdata/fixtures"

func prepareTestDatabase() *sql.DB {
	db, err := ConnectionDB()
	if err != nil {
		log.Fatal(err)
	}
	err = testfixtures.LoadFixtures(FixturesPath, db, &testfixtures.MySQLHelper{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func ConnectionDB() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, "fix_test")
	return sql.Open("mysql", dsn)
}
