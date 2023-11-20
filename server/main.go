package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shinya0226/kensyu/handler"
)

func main() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	e.GET("/", handler.Hello)

	e.POST("/login", handler.Login)

	// サーバーをポート番号8080で起動
	e.Logger.Fatal(e.Start(":8080"))

}