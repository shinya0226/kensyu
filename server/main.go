package main

import (
	"log"

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

	db, err := mysql.ConnectionDB()
	if err != nil {
		log.Fatal(err)
	}
	userRepo := mysql.NewUserRepository(db)
	loginUsecase := usecase.NewLoginUsecase(userRepo)

	// ルートを設定
	e.POST("/login", handler.Login(loginUsecase))

	// サーバーをポート番号8080で起動
	e.Logger.Fatal(e.Start(":8080"))
}
