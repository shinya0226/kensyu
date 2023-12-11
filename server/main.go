package main

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shinya0226/kensyu/entity"
	"github.com/shinya0226/kensyu/handler"
	"github.com/shinya0226/kensyu/infra/mysql"
	"github.com/shinya0226/kensyu/usecase"
)

func main() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := mysql.ConnectionDB()

	userRepo := mysql.NewUserRepository(db)
	loginUsecase := usecase.NewLoginUsecase(userRepo)

	//ログイン処理
	e.POST("/login", Login(loginUsecase))
	// 初期画面
	e.GET("/", handler.Hello)

	//アカウント一覧取得処理
	r := e.Group("/restricted")

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwt.MapClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}
	r.Use(echojwt.WithConfig(config))
	//JWT認証
	r.GET("", restricted)                           //http://localhost:8080/restricted
	r.GET("/accounts/:page", handler.GetAccounts()) // http://localhost:8080/restricted/accounts/1

	// サーバーをポート番号8080で起動
	e.Logger.Fatal(e.Start(":8080"))

}

type AdminFormat struct {
	IsAdmin int `json:"isAdmin"`
}

func restricted(c echo.Context) error {
	// JWT認証
	token := c.Get("user").(*jwt.Token)
	usecase.VerifyToken(token.Raw)

	//isAdmin認証
	if logfo.IsAdmin != 1 {
		return c.String(http.StatusBadRequest, "isAdmin認証NG")
	}

	return c.String(http.StatusOK, "認証OK")
}

// ログイン処理（機能）
func Login(u usecase.ILoginUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		return LoginWithUsecase(u, c)
	}
}

type LoginFormat struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	IsAdmin      int    `json:"isAdmin"`
	Access_token string `json:"access_token"`
}

var logfo LoginFormat

// ログイン処理（詳細）
func LoginWithUsecase(u usecase.ILoginUsecase, c echo.Context) error {
	eu := new(entity.User)
	// logfo := LoginFormat{}

	if err := c.Bind(eu); err != nil {
		return err
	}
	//Loginの出力をmessageに格納
	response, err := u.Login(*eu)
	if err != nil {
		return err
	}
	//formatに追加
	logfo.Email = response.Email
	logfo.Name = response.Name
	logfo.IsAdmin = response.IsAdmin
	logfo.Access_token = response.Access_token

	return c.JSON(http.StatusOK, logfo) //structに詰める
}
