package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/shinya0226/kensyu/entity"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) entity.IUserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) FindSingleRow(email string) (entity.User, error) {
	u := entity.User{}
	// todo 以下でSQL インジェクションが発生しうるかを調査してください
	if err := ur.db.QueryRow("SELECT Email,Password FROM user WHERE Email = ?", email).
		Scan(&u.Email, &u.Password); err != nil {
		//Emailが合致しないとき
		return u, err
	}
	//Emailが合致するとき
	return u, nil
}

// DBに接続
func ConnectionDB() *sql.DB {
	//環境変数の設定
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB")

	//DB読み込み
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUser, DBPass, DBHost, DBPort, DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(http.StatusInternalServerError)
		log.Fatal(err)
	}
	return db
}
