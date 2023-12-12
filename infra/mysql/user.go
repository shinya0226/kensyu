package mysql

import (
	"database/sql"
	"fmt"
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
	dbTable := os.Getenv("DB_TABLE")
	table := fmt.Sprintf("SELECT * FROM '%s' where Email = ?", dbTable)
	if err := ur.db.QueryRow(table, email).
		Scan(&u.Email, &u.Password, &u.Name, &u.IsAdmin); err != nil {
		//　Emailが合致しないとき
		return u, err
	}
	//　Emailが合致するとき
	return u, nil
}
