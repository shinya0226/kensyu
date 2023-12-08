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
	Email        string `json:"email"`
	Name         string `json:"name"`
	IsAdmin      int    `json:"isAdmin"`
	Access_token string `json:"access_token"`
}

func (u *loginUsecase) Login(e entity.User) (LoginFormat, error) {
	//該当するユーザーを抽出（found）
	found, err := u.repo.FindSingleRow(e.Email)

	//出力の型を定義
	logfo := LoginFormat{}

	logfo.Email = found.Email
	logfo.Name = found.Name
	logfo.IsAdmin = found.IsAdmin

	if err != nil {
		return logfo, err
	}
	if e.Password != found.Password {
		return logfo, err
	}
	//JWTの作成
	jwt_message, err := CreateToken(e.Email)
	//出力の型を定義
	logfo.Access_token = jwt_message

	return logfo, err

}
