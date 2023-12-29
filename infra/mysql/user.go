package mysql

import (
	"database/sql"
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
	table := os.Getenv("DB_TABLE")
	if err := ur.db.QueryRow("SELECT * FROM"+" "+table+" "+"WHERE Email = ?", email).
		Scan(&u.Email, &u.Password, &u.Name, &u.IsAdmin); err != nil {
		//　Emailが合致しないとき
		return u, err
	}
	//　Emailが合致するとき
	return u, nil
}
