package handler

import (
	"net/http"

	"github.com/shinya0226/kensyu/usecase"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

// 初期画面の表示
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "お仕事おつかれ様")
}

// ログイン処理
func Login(u usecase.LoginUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		return LoginWithUsecase(u, c)
	}
}

func LoginWithUsecase(u usecase.LoginUsecase, c echo.Context) {
	return c.String(http.StatusOK, "お仕事おつかれ様")
}
