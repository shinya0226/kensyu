package usecase

import "github.com/shinya0226/kensyu/entity"

type loginUsecase struct {
	repo entity.IUserRepository
}

func NewLoginUsecase(repo entity.IUserRepository) ILoginUsecase {
	return &loginUsecase{repo: repo}
}

type ILoginUsecase interface {
	Login(e entity.User) (string, error)
}

func (u *loginUsecase) Login(e entity.User) (string, error) {
	//該当するユーザーを抽出（found）
	found, err := u.repo.FindSingleRow(e.Email)

	if err != nil {
		return "", err
	}
	//DBのパスワードのハッシュ化
	pass, err := HashPassword(e.Password)
	if err != nil {
		return "", err
	}
	//パスワードの比較
	if ans := VerifyPassword(pass, found.Password); ans != nil {
		return "", err
	}
	//JWTの作成
	message, err := CreateToken(e.Email)
	return message, err

}
