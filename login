package main

import (
	"database/sql"
	"fmt"
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
	e.GET("/", Hello)               //ホーム
	e.POST("/login", Login)         //ログイン認証
	e.GET("/accounts", GetAccounts) //アカウント一覧取得

	// サーバーをポート番号8080で起動
	e.Logger.Fatal(e.Start(":8080"))

}

// 初期画面の表示
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "ようこそ！")
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

// 行データを取得
func GetRows(db *sql.DB) *sql.Rows {
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		fmt.Println(http.StatusInternalServerError)
	}
	return rows
}

// ログイン処理
func Login(c echo.Context) error {
	//DB設定
	db := ConnectionDB()
	defer db.Close()
	rows := GetRows(db)

	test := User{}    //空データ作成
	var result []User //すべてのデータを挿入

	for rows.Next() {
		error := rows.Scan(&test.Email, &test.Password, &test.Name, &test.IsAdmin)
		if error != nil {
			fmt.Println(http.StatusInternalServerError)
		} else {
			result = append(result, test)
		}
	}

	//fmt.Println(result) //パスワードを含む配列

	u := new(User) //入力
	if err := c.Bind(u); err != nil {
		fmt.Println(http.StatusBadRequest)
		return err
	}

	for i := 0; i < len(result); i++ {
		pass := result[i].Password
		pass, _ = HashPassword(pass) //DBのパスワードのハッシュ化
		ans := VerifyPassword(pass, u.Password)
		//Emailが合致したとき
		if result[i].Email == u.Email {
			//Passwordが合致しないとき
			if ans != nil {
				fmt.Println(http.StatusNotFound)
				fmt.Println("password error")
				fmt.Println("email:" + u.Email)

				return c.String(http.StatusCreated, "ログイン失敗")
			} else {
				fmt.Println(http.StatusOK)
				fmt.Println("Login OK")
				fmt.Println("email:" + u.Email)
				fmt.Println("name:" + result[i].Name)
				fmt.Print("isAdmin:")
				fmt.Println(u.IsAdmin)

				//JWTの作成
				message, _ := CreateToken(u.Email)
				fmt.Println("access_token:" + message) //アクセストークンの表示

				return c.String(http.StatusCreated, "ログイン成功") //ユーザー画面表示

			}
			//Emailが合致しないとき
		} else {
			fmt.Println(http.StatusNotFound)
			fmt.Println("email error")
			fmt.Println("email:" + u.Email)

			return c.String(http.StatusCreated, "ログイン失敗") //ユーザー画面表示
		}
	}

	return c.JSON(http.StatusCreated, u)
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

// アカウント一覧取得
func GetAccounts(c echo.Context) error {
	return nil
}
