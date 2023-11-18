package handler

import (
	"fmt"
	"net/http"

	"github.com/shinya0226/kensyu/entity"
	"github.com/shinya0226/kensyu/usecase"

	"github.com/shinya0226/kensyu/infra/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	usecase usecase.IUserUsecase
}

func NewUserHandler(uu usecase.IUserUsecase) *userHandler {
	return &userHandler{
		usecase: uu,
	}
}

func (uh userHandler) FindSingleRow() echo.HandlerFunc {
	return func(c echo.Context) error {
		//DB設定
		db := mysql.ConnectionDB()
		//入力処理
		u := new(entity.User)
		if err := c.Bind(u); err != nil {
			return err
		}
		u, err := uh.usecase.FindSingleRow(db, u.Email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		pass := u.Password
		pass, err = usecase.HashPassword(u.Password) //DBのパスワードのハッシュ化

		if err != nil {
			return err
		}

		if ans := usecase.VerifyPassword(pass, u.Password); ans != nil {
			//Passwordが合致しないとき
			fmt.Println("password error")
			fmt.Println("email:" + u.Email)
			return c.JSON(http.StatusBadRequest, u)

		} else {
			//Email,Passwordが合致するとき
			fmt.Println("email:" + u.Email)
			fmt.Println("name:" + u.Name)
			fmt.Print("isAdmin:")
			fmt.Println(u.IsAdmin)
			//JWTの作成
			message, _ := usecase.CreateToken(u.Email)
			fmt.Println("access_token:" + message) //アクセストークンの表示

			return c.JSON(http.StatusCreated, u)
		}
	}
}

// 初期画面の表示
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "お仕事おつかれ様")
}
