package handler

import (
	"fmt"

	"net/http"

	"github.com/shinya0226/kensyu/infra/mysql"
	"github.com/shinya0226/kensyu/usecase"

	"github.com/shinya0226/kensyu/entity"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

// 初期画面の表示
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "お仕事おつかれ様")
}

// ログイン処理
func Login(u usecase.LoginUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		return LoginWithUsecase(u,c)
	}
}

func LoginWithUsecase(u,c){
	
}

	//DB設定
	fmt.Println("DB接続")
	db := mysql.ConnectionDB()
	fmt.Println(db)
	fmt.Println("DB接続完了")

	//入力処理
	u := new(entity.User)
	if err := c.Bind(u); err != nil {
		fmt.Println(err)
		// return err
		return c.String(http.StatusBadRequest, "ログイン失敗")
	}

	var UserRepository IUserRepository
	UserRepository = userRepository{}

	pass := u.Password
	pass, err := usecase.HashPassword(pass) //DBのパスワードのハッシュ化
	if err != nil {
		return c.String(http.StatusBadRequest, "ログイン失敗")
	}

	if ans := usecase.VerifyPassword(pass, u.Password); ans != nil {
		fmt.Println(ans)
		//Passwordが合致しないとき
		return c.String(http.StatusBadRequest, "ログイン失敗")

	} else {
		//JWTの作成
		message, _ := usecase.CreateToken(email)
		fmt.Println("access_token:" + message) //アクセストークンの表示

		// return c.JSON(http.StatusCreated, u)
		return c.String(http.StatusCreated, "ログイン成功")
	}

}
