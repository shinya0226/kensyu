package handler

import (
	"net/http"

	"github.com/shinya0226/kensyu/entity"
	"github.com/shinya0226/kensyu/usecase"

	"github.com/labstack/echo/v4"
)

// ログイン処理（機能）
func Login(u usecase.ILoginUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		return loginWithUsecase(u, c)
	}
}

// ログイン処理（詳細）
func loginWithUsecase(u usecase.ILoginUsecase, c echo.Context) error {
	eu := new(entity.User)
	if err := c.Bind(eu); err != nil {
		return err
	}
	//　Loginの出力をmessageに格納
	if eu.Email == "" || eu.Name == "" || eu.Password == "" {
		return c.String(http.StatusNotFound, "入力値は見つかりません")
	}
	message, err := u.Login(*eu)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, message) //　structに詰める
}
