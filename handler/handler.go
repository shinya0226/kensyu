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

// // 該当するデータを取得
// func FindSingleRow(db *sql.DB, Email string) (string, string, string, int) {
// 	u := entity.User{}
// 	// todo 以下でSQL インジェクションが発生しうるかを調査してください
// 	if err := db.QueryRow("SELECT Email,Password FROM user WHERE Email = ?", Email).
// 		Scan(u.Email, u.Password); err != nil {
// 		//Emailが合致しないとき
// 		fmt.Println("エラー")
// 		fmt.Println(http.StatusNotFound)
// 	}
// 	//Emailが合致するとき
// 	return u.Email, u.Name, u.Password, u.IsAdmin
// }

type User struct {
	Email    string `json:"Email" form:"Email"`
	Password string `json:"Password" form:"Password"`
	Name     string `json:"Name"`
	IsAdmin  int    `json:"IsAdmin"`
}

// 該当するデータを取得
func (user User) FindSingleRow(db *sql.DB, Email string) User {
	//u := &entity.User{}
	// todo 以下でSQL インジェクションが発生しうるかを調査してください
	if err := db.QueryRow("SELECT Email,Password FROM user WHERE Email = ?", Email).
		Scan(&Email, &Password); err != nil {
		//Emailが合致しないとき
		fmt.Println(user)
		fmt.Println(http.StatusNotFound)
	}
	//Emailが合致するとき
	return user
}

// 初期画面の表示
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "お仕事おつかれ様")
}

// ログイン処理
func Login(c echo.Context) error {
	//DB設定
	db := mysql.ConnectionDB()

	//入力処理
	u := new(entity.User)
	if err := c.Bind(u); err != nil {
		fmt.Println(err)
		// return err
		return c.String(http.StatusBadRequest, "ログイン失敗")
	}
	fmt.Println("チェック1")
	fmt.Println(u)
	fmt.Println(u.Email)
	fmt.Println(u.Password)
	fmt.Println(u.Name)
	fmt.Println(u.IsAdmin)
	fmt.Println("チェック")

	// todo 以下のコードのうち、echo に関係ない部分は、 usecase/ に移動する

	//該当するユーザーの情報を抽出
	// email, name, password, isadmin := FindSingleRow(db, u.Email)
	user = u.FindSingleRow(db, u.Email)
	fmt.Println("チェック2")
	fmt.Println(*u)
	fmt.Println(*u)

	fmt.Println(*&u.Email)
	fmt.Println(u.Email)
	fmt.Println(u.Name)
	fmt.Println(u.Password)
	fmt.Println(u.IsAdmin)
	fmt.Println("完了")

	pass := u.Password
	pass, err := usecase.HashPassword(pass) //DBのパスワードのハッシュ化
	if err != nil {
		// fmt.Println(err)
		// fmt.Println(http.StatusInternalServerError)
		// return err
		return c.String(http.StatusBadRequest, "ログイン失敗")
	}

	// if ans := usecase.VerifyPassword(pass, u.Password); ans != nil {
	// 	fmt.Println(ans)
	// 	//Passwordが合致しないとき
	// 	// fmt.Println(http.StatusNotFound)
	// 	// fmt.Println(err)
	// 	fmt.Println(u.Email)
	// 	fmt.Println(email)
	// 	fmt.Println(pass)
	// 	fmt.Println(u.Password)
	// 	fmt.Println("password error")
	// 	fmt.Println("email:" + email)
	// 	// return c.JSON(http.StatusBadRequest, u)
	// 	return c.String(http.StatusBadRequest, "ログイン失敗")

	// } else {
	// 	//Email,Passwordが合致するとき
	// 	// fmt.Println(http.StatusOK)
	// 	// fmt.Println("Login OK")
	// 	fmt.Println("email:" + email)
	// 	fmt.Println("name:" + name)
	// 	fmt.Print("isAdmin:")
	// 	fmt.Println(isadmin)
	// 	//JWTの作成
	// 	message, _ := usecase.CreateToken(email)
	// 	fmt.Println("access_token:" + message) //アクセストークンの表示

	// 	// return c.JSON(http.StatusCreated, u)
	return c.String(http.StatusCreated, "ログイン成功")

}
