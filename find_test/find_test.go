package find_test

import (
	"database/sql"
	"fmt"

	"github.com/shinya0226/kensyu/entity"
	"github.com/shinya0226/kensyu/usecase"
)

type userRepository struct {
	db *sql.DB
}

func (ur *userRepository) TestFindSingleRow(email string) (entity.User, error) {
	u := entity.User{}
	// todo 以下でSQL インジェクションが発生しうるかを調査してください
	if err := ur.db.QueryRow("SELECT * FROM user WHERE Email = ?", email).
		Scan(&u.Email, &u.Password, &u.Name, &u.IsAdmin); err != nil {
		//Emailが合致しないとき
		return u, err
	}
	//Emailが合致するとき
	return u, nil
}

type loginUsecase struct {
	repo entity.IUserRepository
}

func (u *loginUsecase) Login(e entity.User) (string, error) {
	//該当するユーザーを抽出（found）
	found, err := u.repo.FindSingleRow(e.Email)

	if err != nil {
		fmt.Println("エラー1")
		return "", err
	}
	//DBのパスワードのハッシュ化
	pass, err := usecase.HashPassword(e.Password)
	if err != nil {
		fmt.Println("エラー2")
		return "", err
	}
	//パスワードの比較
	if ans := usecase.VerifyPassword(pass, found.Password); ans != nil {
		fmt.Println("エラー3")
		return "", err
	}
	//JWTの作成
	message, err := usecase.CreateToken(e.Email)
	fmt.Println("エラー4")
	return message, err

}
