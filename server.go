package main

import (
	"net/http"

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
	e.GET("/", Hello)

	h := handler.Handler{} //TODO: DBをいれる？
	e.POST("/login", func(ctx echo.Context) error {
		return h.Login(ctx)
	})

	// サーバーをポート番号8080で起動
	e.Logger.Fatal(e.Start(":8080"))

}

// 初期画面の表示
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "お仕事おつかれ様")
}
