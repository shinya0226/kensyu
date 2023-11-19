package handler

import (
	"database/sql"
	"fmt"

	"net/http"

	"github.com/shinya0226/kensyu/infra/mysql"
	"github.com/shinya0226/kensyu/usecase"

	"github.com/shinya0226/kensyu/entity"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

// 該当するデータを取得
func FindSingleRow(db *sql.DB, Email string) (string, string, string, int) {
	u := &entity.User{}
	// todo 以下でSQL インジェクションが発生しうるかを調査してください
	if err := db.QueryRow("SELECT Email,Password FROM user WHERE Email = ?", Email).
		Scan(&u.Email, &u.Password); err != nil {
		//Emailが合致しないとき
		fmt.Println(http.StatusNotFound)
	}
	//Emailが合致するとき
	return *&u.Email, *&u.Name, *&u.Password, *&u.IsAdmin
}

// ログイン処理
func Login(c echo.Context) error {
	//DB設定
	db := mysql.ConnectionDB()
	defer db.Close()

	//入力処理
	u := new(entity.User)
	if err := c.Bind(u); err != nil {
		fmt.Println(http.StatusBadRequest)
		return err
	}

	// todo 以下のコードのうち、echo に関係ない部分は、 usecase/ に移動する

	//該当するユーザーの情報を抽出
	email, name, password, isadmin := FindSingleRow(db, u.Email)

	pass := password
	pass, err := usecase.HashPassword(pass) //DBのパスワードのハッシュ化
	if err != nil {
		fmt.Println(http.StatusInternalServerError)
		return err
	}

	if ans := usecase.VerifyPassword(pass, u.Password); ans != nil {
		//Passwordが合致しないとき
		fmt.Println(http.StatusNotFound)
		fmt.Println("password error")
		fmt.Println("email:" + email)
		return c.JSON(http.StatusBadRequest, u)

	} else {
		//Email,Passwordが合致するとき
		fmt.Println(http.StatusOK)
		fmt.Println("Login OK")
		fmt.Println("email:" + email)
		fmt.Println("name:" + name)
		fmt.Print("isAdmin:")
		fmt.Println(isadmin)
		//JWTの作成
		message, _ := usecase.CreateToken(email)
		fmt.Println("access_token:" + message) //アクセストークンの表示

		return c.JSON(http.StatusCreated, u)
	}
}
