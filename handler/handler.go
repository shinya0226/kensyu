package handler

import (
	"net/http"

	"github.com/shinya0226/kensyu/entity"
	"github.com/shinya0226/kensyu/usecase"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

// 初期画面の表示
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "お仕事おつかれ様")
}

// ログイン処理（機能）
func Login(u usecase.ILoginUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		return LoginWithUsecase(u, c)
	}
}

type LoginFormat struct {
	email        string `json:"email"`
	name         string `json:"Name"`
	isAdmin      int    `json:"isAdmin"`
	access_token string `json:"access_token"`
}

// ログイン処理（詳細）
func LoginWithUsecase(u usecase.ILoginUsecase, c echo.Context) error {
	eu := new(entity.User)
	logfo := LoginFormat{}

	if err := c.Bind(eu); err != nil {
		return err
	}
	message, err := u.Login(*eu)
	if err != nil {
		return err
	}
	//formatに追加
	logfo.email = eu.Email
	logfo.name = eu.Name
	logfo.isAdmin = eu.IsAdmin
	logfo.access_token = message

	return c.JSON(http.StatusOK, logfo) //structに詰める
}
