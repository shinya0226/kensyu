package handler

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/shinya0226/kensyu/entity"
	"github.com/shinya0226/kensyu/infra/mysql"
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
	Email        string `json:"email"`
	Name         string `json:"name"`
	IsAdmin      int    `json:"isAdmin"`
	Access_token string `json:"access_token"`
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	// jwt.RegisteredClaims
}

var logfo LoginFormat

// ログイン処理（詳細）
func LoginWithUsecase(u usecase.ILoginUsecase, c echo.Context) error {
	eu := new(entity.User)
	// logfo := LoginFormat{}

	if err := c.Bind(eu); err != nil {
		return err
	}
	//Loginの出力をmessageに格納
	response, err := u.Login(*eu)
	if err != nil {
		return err
	}
	//formatに追加
	logfo.Email = response.Email
	logfo.Name = response.Name
	logfo.IsAdmin = response.IsAdmin
	logfo.Access_token = response.Access_token

	return c.JSON(http.StatusOK, logfo) //structに詰める
}

// アカウント一覧取得
func GetAccounts() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := mysql.ConnectionDB()
		//JWT認証
		user := c.Get("user").(*jwt.Token)
		_, err := usecase.VerifyToken(user.Raw)
		if err != nil {
			return err
		}

		post := entity.User{}
		posts := []*entity.User{}
		//request page
		page := c.Param("page")
		//int型に変換
		var i int
		i, _ = strconv.Atoi(page)
		//読み込み開始のページの定義
		var page_first int
		page_first = (i - 1) * 5

		rows, err := db.Query("select * from user LIMIT ?,5", page_first)
		if err != nil {
			return err
		}
		for rows.Next() {
			if err := rows.Scan(&post.Email, &post.Password, &post.Name, &post.IsAdmin); err != nil {
				return err
			}
			posts = append(posts, &entity.User{Email: post.Email, Name: post.Name, IsAdmin: post.IsAdmin})
		}
		return c.JSON(http.StatusOK, posts)
	}
}
