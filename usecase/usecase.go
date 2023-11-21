package usecase

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/shinya0226/kensyu/entity"
	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase struct {
	repo entity.IUserRepository
}

func (u *LoginUsecase) Login(e entity.User) (string, error) {
	found, err := u.repo.FindSingleRow(e.Email)
	if err != nil {
		return "", err
	}

	pass, err := HashPassword(e.Password) //DBのパスワードのハッシュ化
	if err != nil {
		return "", err
	}

	if ans := VerifyPassword(pass, found.Password); ans != nil {
		//Passwordが合致しないとき
		return "", err

	}
	//JWTの作成
	message, err := CreateToken(e.Email)
	return message, err

}

// パスワードの暗号化（DBからパスワードを取り出す時に使用）
func HashPassword(rawPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// 暗号化されたパスワードとユーザーが入力したパスワードの比較
func VerifyPassword(hashedPassword string, entryPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(entryPassword))
	return err
}

// JWTの発行
func CreateToken(email string) (string, error) {
	//tokenの作成
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	//Claimsの設定
	token.Claims = jwt.MapClaims{
		"user": email,
		"exp":  time.Now().Add(time.Hour * 1).Unix(), //1時間の有効期限を設定
	}
	//署名
	JWT_SECRET := os.Getenv("JWT_SECRET")
	var secretKey = JWT_SECRET
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return string(http.StatusInternalServerError), err

	}
	return tokenString, err

}

// JWTの検証
func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return token, err
	}
	return token, err
}
