package usecase

import (
	"github.com/shinya0226/kensyu/entity"
)

// ログイン処理
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
	Email       string `json:"email"`
	Name        string `json:"name"`
	IsAdmin     int    `json:"isAdmin"`
	AccessToken string `json:"access_token"`
}

var logfo LoginFormat

func (u *loginUsecase) Login(e entity.User) (LoginFormat, error) {
	//　該当するユーザーを抽出（found）
	found, err := u.repo.FindSingleRow(e.Email)
	//　Emailの合致確認
	if err != nil {
		return LoginFormat{}, err
	}

	logfo.Email = found.Email

	//　Passwordの合致確認
	err = verifyPassword(found.Password, e.Password)
	if err != nil {
		return LoginFormat{}, err
	}
	logfo.Name = found.Name
	logfo.IsAdmin = found.IsAdmin
	//　JWTの作成
	jwtMessage := createToken(e.Email)
	//　出力の型を定義
	logfo.AccessToken = jwtMessage
	return logfo, nil
}
