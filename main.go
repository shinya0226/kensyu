package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	e.GET("/", Hello)
	e.POST("/login", Login)

	// サーバーをポート番号8080で起動
	e.Logger.Fatal(e.Start(":8080"))

}

// 初期画面の表示
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "お仕事おつかれ様！")
}

// データの型を定義
type User struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Name     string `json:"Name"`
	IsAdmin  int    `json:"IsAdmin"`
}

// DBに接続
func ConnectionDB() *sql.DB {
	dsn := "root:Shinya0023@tcp(127.0.0.1:3306)/yamamoto?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(http.StatusInternalServerError)
	}
	return db
}

// 該当するデータを取得
func FindSingleRow(db *sql.DB, Email string) (*sql.DB, string, string, string, int) {
	u := &User{}
	if err := db.QueryRow("SELECT * FROM user WHERE Email = ?", Email).
		Scan(&u.Email, &u.Password, &u.Name, &u.IsAdmin); err != nil {
		//Emailが合致しないとき
		fmt.Println(http.StatusNotFound)
		fmt.Println("email error")
		fmt.Println("email:" + Email)
		log.Fatal(err)
	}
	//Emailが合致するとき
	return db, *&u.Email, *&u.Name, *&u.Password, *&u.IsAdmin
}

// ログイン処理
func Login(c echo.Context) error {
	//DB設定
	db := ConnectionDB()
	defer db.Close()

	//入力処理
	u := new(User)
	if err := c.Bind(u); err != nil {
		fmt.Println(http.StatusBadRequest)
		return err
	}

	//該当するユーザーの情報を抽出
	db, email, name, password, isadmin := FindSingleRow(db, u.Email)

	pass := password
	pass, _ = HashPassword(pass) //DBのPasswordのハッシュ化

	if ans := VerifyPassword(pass, u.Password); ans != nil {
		//Passwordが合致しないとき
		fmt.Println(http.StatusNotFound)
		fmt.Println("password error")
		fmt.Println("email:" + email)
		return c.String(http.StatusCreated, "ログイン失敗")

	} else {
		//Email,Passwordが合致するとき
		fmt.Println(http.StatusOK)
		fmt.Println("Login OK")
		fmt.Println("email:" + email)
		fmt.Println("name:" + name)
		fmt.Print("isAdmin:")
		fmt.Println(isadmin)
		//JWTの作成
		message, _ := CreateToken(email)
		fmt.Println("access_token:" + message) //アクセストークンの表示

		return c.String(http.StatusCreated, "ログイン成功") //ユーザー画面表示
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
