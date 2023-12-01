package mysql

import (
	"database/sql"

	"github.com/shinya0226/kensyu/entity"
)

// ログイン処理
type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) entity.IUserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) FindSingleRow(email string) (entity.User, error) {
	u := entity.User{}
	// todo 以下でSQL インジェクションが発生しうるかを調査してください
	if err := ur.db.QueryRow("SELECT * FROM user WHERE Email = ?", email).
		Scan(&u.Email, &u.Password, &u.Name, &u.IsAdmin); err != nil {
		//Emailが合致しないとき
		return u, err
	} else {
		//Emailが合致するとき
		return u, nil
	}
}
