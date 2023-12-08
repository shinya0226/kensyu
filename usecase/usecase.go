package usecase

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

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

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

// JWTの発行
func CreateToken(email string) (string, error) {
	//tokenの作成
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	//Claimsの設定
	token.Claims = &jwt.MapClaims{
		"name": email,
		"user": email,
		"exp":  time.Now().Add(time.Hour * 1).Unix(), //1時間の有効期限を設定
	}

	//署名
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err

	}
	return tokenString, err

}

// JWTの検証
func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv(os.Getenv("JWT_SECRET"))), nil
	})
	if err != nil {
		return token, err
	}
	return token, err
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
