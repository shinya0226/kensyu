package main

import (
	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"fmt"
	"github.com/shinya0226/kensyu/infra/mysql"
	"gopkg.in/testfixtures.v1"
	"log"
	"path/filepath"
)

func main() {
	abspath, _ := filepath.Abs(FixturesPath)
	fmt.Println(abspath)
	prepareTestDatabase()
}

const FixturesPath = "testdata/fixtures"

func prepareTestDatabase() *sql.DB {
	db, err := mysql.ConnectionDB()
	if err != nil {
		log.Fatal(err)
	}
	err = testfixtures.LoadFixtures(FixturesPath, db, &testfixtures.MySQLHelper{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
