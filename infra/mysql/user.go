package mysql

import (
	"database/sql"
	"fmt"

	"github.com/joho/godotenv"
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
	//環境設定ファイルの読み込み
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf("読み込み失敗: %v", err)
	}

	if err := ur.db.QueryRow("SELECT * FROM test WHERE Email = ?", email).
		Scan(&u.Email, &u.Password, &u.Name, &u.IsAdmin); err != nil {
		//Emailが合致しないとき
		return u, err
	} else {
		//Emailが合致するとき
		return u, nil
	}
}
