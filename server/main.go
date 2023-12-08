package main

import (
	"github.com/golang-jwt/jwt"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	// 初期画面
	e.GET("/", handler.Hello)
	//ログイン処理
	e.POST("/login", handler.Login(loginUsecase))
	//アカウント一覧取得処理
	// e.POST("/accounts/:page", handler.GetAccounts())
	r := e.Group("/restricted")

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	r.Use(echojwt.WithConfig(config))
	r.GET("", restricted)

	// r.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))
	// r.POST("", handler.Restricted())

	// サーバーをポート番号8080で起動
	e.Logger.Fatal(e.Start(":8080"))

}
