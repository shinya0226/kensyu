package handler

import (
	"fmt"
	"net/http"
	"os"
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
		return c.String(http.StatusNotFound, "入力値は見つかりません")
	}
	message, err := u.Login(*eu)
	if err != nil {
		return c.String(http.StatusNotFound, "Emailは見つかりません")
	}
	return c.JSON(http.StatusOK, message) //　structに詰める
}

type AdminFormat struct {
	IsAdmin int `json:"isAdmin"`
}

func Allowed(c echo.Context) error {
	var logfo LoginFormat
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
		table := os.Getenv("DB_TABLE")
		sql := fmt.Sprintf("select * from" + " " + table + " " + "LIMIT ?,5")
		rows, err := db.Query(sql, pagefirst)
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
		if eu.Email == "" || eu.Password == "" || eu.Name == "" {
			return c.String(http.StatusNotFound, "入力値はありません")
		}
		db := mysql.ConnectionDB()
		defer db.Close()
		table := os.Getenv("DB_TABLE")
		sql := fmt.Sprintf("INSERT INTO" + " " + table + " " + "VALUES(?,?,?,?)")
		ins, err := db.Prepare(sql)
		if err != nil {
			return err
		}
		defer ins.Close()
		pass, err := hashPassword(eu.Password)
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
func hashPassword(rawPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// アカウント削除
func DeleteAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := mysql.ConnectionDB()
		defer db.Close()
		df := new(DeleteFormat)
		if err := c.Bind(df); err != nil {
			return err
		}
		if df.Email == "" {
			return c.String(http.StatusNotFound, "入力値はありません")
		}
		table := os.Getenv("DB_TABLE")
		sql := fmt.Sprintf("DELETE FROM" + " " + table + " " + "WHERE Email = ?")
		del, err := db.Prepare(sql)
		if err != nil {
			return err
		}
		defer del.Close()
		res, err := del.Exec(df.Email)
		if err != nil || res == nil {
			return err
		}
		return c.String(http.StatusOK, "削除完了")
	}
}

type DeleteFormat struct {
	Email string `json:"email"`
}

// アカウント更新
func UpdateAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := mysql.ConnectionDB()
		defer db.Close()
		ua := new(UpdateFormat)
		if err := c.Bind(ua); err != nil {
			return err
		}
		if ua.Email == "" || ua.Password == "" {
			return c.String(http.StatusNotFound, "入力値はありません")
		}
		pass, err := hashPassword(ua.Password)
		if err != nil {
			return err
		}
		table := os.Getenv("DB_TABLE")
		sql := fmt.Sprintf("UPDATE" + " " + table + " " + "SET Password = ?, IsAdmin = ? WHERE Email = ?")
		upd, err := db.Prepare(sql)
		//　upd, err := db.Prepare("UPDATE users SET Password = ?, IsAdmin = ? WHERE Email = ?")
		if err != nil {
			return err
		}
		defer upd.Close()
		res, err := upd.Exec(pass, ua.IsAdmin, ua.Email)
		if err != nil || res == nil {
			return err
		}
		return c.String(http.StatusOK, "更新完了")
	}
}

type UpdateFormat struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  int    `json:"isAdmin"`
}
