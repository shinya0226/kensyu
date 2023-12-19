package usecase

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// 暗号化されたパスワードとユーザーが入力したパスワードの比較
func verifyPassword(hashedPassword string, entryPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(entryPassword))
	return err
}

// JWTの発行
func createToken(email string) (string, error) {
	//　tokenの作成
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	//　Claimsの設定
	token.Claims = jwt.MapClaims{
		"user": email,
		"exp":  time.Now().Add(time.Hour * 1).Unix(), //　1時間の有効期限を設定
	}
	//　署名
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return tokenString, nil
}
