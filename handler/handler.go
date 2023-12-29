package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/shinya0226/kensyu/entity"
	"github.com/shinya0226/kensyu/infra/mysql"
	"github.com/shinya0226/kensyu/usecase"
)

func Login(u usecase.ILoginUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		return loginWithUsecase(u, c)
	}
}

func loginWithUsecase(u usecase.ILoginUsecase, c echo.Context) error {
	eu := new(entity.User)
	if err := c.Bind(eu); err != nil {
		return err
	}
	//　Loginの出力をmessageに格納
	if eu.Email == "" || eu.Password == "" {
		return c.String(http.StatusNotFound, "入力値は見つかりません")
	}
	message, err := u.Login(*eu)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, message) //　structに詰める
}

type AdminFormat struct {
	IsAdmin int `json:"isAdmin"`
}

type LoginFormat struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	IsAdmin     int    `json:"isAdmin"`
	AccessToken string `json:"access_token"`
}

func Allowed(c echo.Context) error {
	var logfo usecase.LoginFormat
	// JWT認証
	token, _ := c.Get("user").(*jwt.Token)
	_, err := usecase.VerifyToken(token.Raw)
	if err != nil {
		return err
	}
	if logfo.IsAdmin != 1 {
		return c.String(http.StatusBadRequest, "isAdmin認証NG")
	}
	return c.String(http.StatusOK, "認証OK")
}

// アカウント一覧取得
func FetchAccounts() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := mysql.ConnectionDB()
		defer db.Close()
		post := entity.User{}
		posts := []*entity.User{}
		//　request page
		page := c.Param("page")
		//　int型に変換
		var i int
		i, _ = strconv.Atoi(page)
		//　読み込み開始のページの定義
		pageFirst := i - 1
		paging := 5
		pagefirst := pageFirst * paging
		rows, err := db.Query("select * from users LIMIT ?,5", pagefirst)
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			err2 := rows.Scan(&post.Email, &post.Password, &post.Name, &post.IsAdmin)
			if err2 != nil {
				return err
			}
			posts = append(posts, &entity.User{Email: post.Email, Name: post.Name, IsAdmin: post.IsAdmin})
		}
		err = rows.Err()
		if err != nil {
			return err
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
		defer db.Close()
		execSequ := fmt.Sprintf("INSERT INTO users VALUES(?,?,?,?)")
		ins, err := db.Prepare(execSequ)
		if err != nil {
			return err
		}
		defer ins.Close()
		pass, err := hashPassword(eu.Password)
		if err != nil {
			return err
		}
		res, err := ins.Exec(eu.Email, pass, eu.Name, eu.IsAdmin)
		if err != nil {
			return err
		}
		if res == nil {
			return err
		}
		return c.String(http.StatusCreated, "作成完了")
	}
}

func hashPassword(rawPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	return string(hashedPassword), err
}
