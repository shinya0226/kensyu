package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"golang.org/x/crypto/bcrypt"
)

type (
	User struct {
		Email    string `json:"Email" form:"Email"`
		Password string `json:"Password" form:"Password"`
	}
	Handler struct {
		DB map[string]*User
	}
)

// DBに接続
func ConnectionDB() *sql.DB {
	dsn := "root:Shinya0023@tcp(127.0.0.1:3306)/yamamoto?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(http.StatusInternalServerError)
		log.Fatal(err)
	}
	return db
}

// 該当するデータを取得
func FindSingleRow(db *sql.DB, Email string) (*sql.DB, string, string) {
	u := &User{}
	if err := db.QueryRow("SELECT Email,Password FROM user WHERE Email = ?", Email).
		Scan(&u.Email, &u.Password); err != nil {
		//Emailが合致しないとき
		fmt.Println(http.StatusNotFound)
		log.Fatal(err)
	}
	//Emailが合致するとき
	return db, *&u.Email, *&u.Password
}

// ログイン処理
func (h *Handler) Login(c echo.Context) error {
	//DB設定
	db := ConnectionDB()
	defer db.Close()

	//入力処理
	u := new(User)
	if err := c.Bind(u); err != nil {
		fmt.Println(http.StatusBadRequest)
		log.Fatal(err)
		return err
	}

	//該当するユーザーの情報を抽出
	db, email, password := FindSingleRow(db, u.Email)

	pass := password
	pass, _ = HashPassword(pass) //DBのパスワードのハッシュ化

	if ans := VerifyPassword(pass, u.Password); ans != nil {
		//Passwordが合致しないとき
		fmt.Println(http.StatusNotFound)
		fmt.Println("password error")
		fmt.Println("email:" + email)
		return c.JSON(http.StatusCreated, u)

	} else {
		//Email,Passwordが合致するとき
		fmt.Println(http.StatusOK)
		fmt.Println("Login OK")
		fmt.Println("email:" + email)
		//JWTの作成
		message, _ := CreateToken(email)
		fmt.Println("access_token:" + message) //アクセストークンの表示

		return c.JSON(http.StatusCreated, u)
	}
}

// パスワードの暗号化（DBからパスワードを取り出す時に使用）
func HashPassword(rawPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// 暗号化されたパスワードとユーザーが入力したパスワードの比較
func VerifyPassword(hashedPassword string, entryPassword string) error {
	//password:=[]byte(hashedPassword)
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
	var secretKey = "secret"
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println(http.StatusInternalServerError)
	}
	return tokenString, nil

}

// JWTの検証
func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		fmt.Println(http.StatusInternalServerError)
		return token, err
	}
	return token, nil
}
