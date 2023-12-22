package handler

import (
	"net/http"
	"strconv"

	"github.com/shinya0226/kensyu/entity"
	"github.com/shinya0226/kensyu/infra/mysql"
	usecase "github.com/shinya0226/kensyu/usecase"

	"github.com/labstack/echo/v4"
)

type AdminFormat struct {
	IsAdmin int `json:"isAdmin"`
}

// ログイン処理（機能）
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
	if eu.Email == "" || eu.Name == "" || eu.Password == "" {
		return c.String(http.StatusNotFound, "入力値は見つかりません")
	}
	message, err := u.Login(*eu)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, message) //　structに詰める
}

// アカウント一覧取得
func GetAccounts() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := mysql.ConnectionDB()
		db.Close()
		post := entity.User{}
		posts := []*entity.User{}
		//　request page
		page := c.Param("page")
		//　int型に変換
		var i int
		i, _ = strconv.Atoi(page)
		//　読み込み開始のページの定義
		pageFirst := (i - 1)
		pageFirst *= 5

		rows, _ := db.Query("select * from users LIMIT ?,5", pageFirst)
		rows.Close()
		err := rows.Err()
		if err != nil {
			return err
		}
		for rows.Next() {
			err := rows.Scan(&post.Email, &post.Password, &post.Name, &post.IsAdmin)
			if err != nil {
				return err
			}
			posts = append(posts, &entity.User{Email: post.Email, Name: post.Name, IsAdmin: post.IsAdmin})
		}
		return c.JSON(http.StatusOK, posts)
	}
}
