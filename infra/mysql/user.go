package mysql

import (
	"database/sql"

	"github.com/shinya0226/kensyu/entity"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) entity.IUserRepository {
	return &userRepository{db: db}
}

// fixtureのファイルパス
const FixturesPathSQL = "../../testdata/fixtures"

func (ur *userRepository) FindSingleRow(email string) (entity.User, error) {
	//　fixture追加
	u := entity.User{}
	db := ConnectionDB()
	db.Close()
	if err := ur.db.QueryRow("SELECT * FROM users where Email = ?", email).
		Scan(&u.Email, &u.Password, &u.Name, &u.IsAdmin); err != nil {
		//　Emailが合致しないとき
		return u, err
	}
	//　Emailが合致するとき
	return u, nil
}
