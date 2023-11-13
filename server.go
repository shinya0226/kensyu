package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	e.GET("/", Hello)

	// サーバーをポート番号8080で起動
	e.Logger.Fatal(e.Start(":8080"))

}

// 初期画面の表示
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "お仕事おつかれ様")
}
