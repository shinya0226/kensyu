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
	Login(e entity.User) (entity.LoginFormat, error)
}

func (u *loginUsecase) Login(e entity.User) (entity.LoginFormat, error) {
	var logfo entity.LoginFormat
	//　該当するユーザーを抽出（found）
	found, err := u.repo.FindSingleRow(e.Email)
	//　Emailの合致確認
	if err != nil {
		return entity.LoginFormat{}, err
	}

	logfo.Email = found.Email

	//　Passwordの合致確認
	err = verifyPassword(found.Password, e.Password)
	if err != nil {
		return entity.LoginFormat{}, err
	}
	logfo.Name = found.Name
	logfo.IsAdmin = found.IsAdmin
	//　JWTの作成
	jwtMessage := createToken(e.Email)
	logfo.AccessToken = jwtMessage
	return logfo, nil
}
