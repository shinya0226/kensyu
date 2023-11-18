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
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) entity.IUserRepository {
	return &userRepository{
		DB: db,
	}
}

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	//環境変数の設定
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB")

	//DB読み込み
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUser, DBPass, DBHost, DBPort, DBName)

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error)
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn

	return sqlHandler
}

// 該当するデータを取得
func (ur *userRepository) FindSingleRow(db *sql.DB, Email string) (*entity.User, error) {
	// todo 以下でSQL インジェクションが発生しうるかを調査してください
	u := entity.User{}
	if err := db.QueryRow("SELECT Email,Password FROM user WHERE Email = ?", Email).
		Scan(&u.Email, &u.Password); err != nil {
		//Emailが合致しないとき
		return &entity.User{}, err
	}
	//Emailが合致するとき
	return &entity.User{}, nil
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
