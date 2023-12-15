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
		return LoginWithUsecase(u, c)
	}
}

type LoginFormat struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	IsAdmin     int    `json:"isAdmin"`
	AccessToken string `json:"access_token"`
}

// ログイン処理（詳細）
func LoginWithUsecase(u usecase.ILoginUsecase, c echo.Context) error {
	eu := new(entity.User)
	logfo := LoginFormat{}

	if err := c.Bind(eu); err != nil {
		return err
	}
	//　Loginの出力をmessageに格納（修正）
	message, err := u.Login(*eu)
	if err != nil {
		return err
	}
	//　formatに追加
	logfo.Email = message.Email
	logfo.Name = message.Name
	logfo.IsAdmin = message.IsAdmin
	logfo.AccessToken = message.AccessToken

	return c.JSON(http.StatusOK, logfo) //　structに詰める
}

func LoginAccount(u usecase.ILoginUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		eu := new(entity.User)
		logfo := LoginFormat{}

		if err := c.Bind(eu); err != nil {
			return err
		}
		//　Loginの出力をmessageに格納（修正）
		message, err := u.Login(*eu)
		if err != nil {
			return err
		}
		//　formatに追加
		logfo.Email = message.Email
		logfo.Name = message.Name
		logfo.IsAdmin = message.IsAdmin
		logfo.AccessToken = message.AccessToken

		return c.JSON(http.StatusOK, logfo) //　structに詰める
	}
}
