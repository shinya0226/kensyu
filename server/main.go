package main

import (
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
	e.GET("/accounts", handler.GetAccounts())

	// サーバーをポート番号8080で起動
	e.Logger.Fatal(e.Start(":8080"))

}
