package handler

import (
	"net/http"
	"strconv"

	"github.com/shinya0226/kensyu/entity"
	"github.com/shinya0226/kensyu/infra/mysql"
	usecase "github.com/shinya0226/kensyu/usecase"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

// 初期画面の表示
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "お仕事おつかれ様")
}

type AdminFormat struct {
	IsAdmin int `json:"isAdmin"`
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

var logfo LoginFormat

// ログイン処理（詳細）
func LoginWithUsecase(u usecase.ILoginUsecase, c echo.Context) error {
	eu := new(entity.User)
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
func FetchAccounts() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := mysql.ConnectionDB()

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

// アカウント作成
func CreateAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		eu := new(entity.User)
		if err := c.Bind(eu); err != nil {
			return err
		}
		db := mysql.ConnectionDB()
		ins, err := db.Prepare("INSERT INTO user VALUES(?,?,?,?)")
		if err != nil {
			return err
		}
		pass, err := usecase.HashPassword(eu.Password)
		if err != nil {
			return err
		}
		res, err := ins.Exec(eu.Email, pass, eu.Name, eu.IsAdmin)
		if err != nil || res == nil {
			return err
		}
		return c.String(http.StatusCreated, "作成完了")
	}
}

type DeleteFormat struct {
	Email string `json:"email"`
}

// アカウント削除
func DeleteAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := mysql.ConnectionDB()
		df := new(DeleteFormat)
		if err := c.Bind(df); err != nil {
			return err
		}

		del, err := db.Prepare("DELETE FROM user WHERE Email = ?")
		if err != nil {
			return err
		}
		res, err := del.Exec(df.Email)
		if err != nil || res == nil {
			return err
		}
		return c.String(http.StatusOK, "削除完了")

	}
}

type UpdateFormat struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  int    `json:"isAdmin"`
}

// アカウント更新
func UpdateAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := mysql.ConnectionDB()
		ua := new(UpdateFormat)
		if err := c.Bind(ua); err != nil {
			return err
		}
		pass, err := usecase.HashPassword(ua.Password)
		if err != nil {
			return err
		}
		upd, err := db.Prepare("UPDATE user SET Password = ?, IsAdmin = ? WHERE Email = ?")
		if err != nil {
			return err
		}
		res, err := upd.Exec(pass, ua.IsAdmin, ua.Email)
		if err != nil || res == nil {
			return err
		}
		return c.String(http.StatusOK, "更新完了")
	}
}
