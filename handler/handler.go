package handler

import (
	"net/http"

	"github.com/shinya0226/kensyu/entity"
	"github.com/shinya0226/kensyu/usecase"

	"github.com/labstack/echo/v4"
)

func Login(u usecase.ILoginUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		return loginWithUsecase(u, c)
	}
}

type LoginFormat struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	IsAdmin     int    `json:"isAdmin"`
	AccessToken string `json:"access_token"`
}

// ログイン処理（詳細）
func loginWithUsecase(u usecase.ILoginUsecase, c echo.Context) error {
	eu := new(entity.User)
	if err := c.Bind(eu); err != nil {
		return err
	}
	//　Loginの出力をmessageに格納
	if eu.Email == "" || eu.Password == "" {
		err := entity.Err{Error: "input is nil and not found"}
		return c.JSON(http.StatusNotFound, err)
	}
	message, err := u.Login(*eu)
	if err != nil {
		err := entity.Err{Error: "input is nil and not found"}
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, message) //　structに詰める
}
