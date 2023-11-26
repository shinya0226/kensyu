package usecase

import (
	"github.com/shinya0226/kensyu/entity"
)

type loginUsecase struct {
	repo entity.IUserRepository
}

func NewLoginUsecase(repo entity.IUserRepository) ILoginUsecase {
	return &loginUsecase{repo: repo}
}

type ILoginUsecase interface {
	Login(e entity.User) (LoginFormat, error)
}

type LoginFormat struct {
	email        string
	name         string
	isAdmin      int
	access_token string
}

func (u *loginUsecase) Login(e entity.User) (LoginFormat, error) {
	//該当するユーザーを抽出（found）
	found, err := u.repo.FindSingleRow(e.Email)

	//出力の型を定義
	logfo := LoginFormat{}

	logfo.email = found.Email
	logfo.name = found.Name
	logfo.isAdmin = found.IsAdmin

	if err != nil {
		return logfo, err
	}
	//DBのパスワードのハッシュ化
	pass, err := HashPassword(e.Password)
	if err != nil {
		return logfo, err
	}
	//パスワードの比較
	if ans := VerifyPassword(pass, found.Password); ans != nil {
		return logfo, err
	}
	//JWTの作成
	jwt_message, err := CreateToken(e.Email)
	//出力の型を定義
	logfo.access_token = jwt_message

	return logfo, err

}
