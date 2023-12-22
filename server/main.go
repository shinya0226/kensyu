package main

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shinya0226/kensyu/handler"
	"github.com/shinya0226/kensyu/infra/mysql"
	. "github.com/shinya0226/kensyu/usecase"
)

func main() {
	// インスタンスを作成
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := mysql.ConnectionDB()
	userRepo := mysql.NewUserRepository(db)
	loginUsecase := NewLoginUsecase(userRepo)

	e.POST("/login", handler.Login(loginUsecase))
	r := e.Group("/restricted")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwt.MapClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}
	r.Use(echojwt.WithConfig(config))
	//　JWT認証
	r.GET("", restricted)                           //　http://localhost:8080/restricted
	r.GET("/accounts/:page", handler.GetAccounts()) //　http://localhost:8080/restricted/accounts/1
	e.Logger.Fatal(e.Start(":8080"))
}

type AdminFormat struct {
	IsAdmin int `json:"isAdmin"`
}

func restricted(c echo.Context) error {
	// JWT認証
	token := c.Get("user").(*jwt.Token)
	_, err := VerifyToken(token.Raw)
	if err != nil {
		return err
	}
	if Logfo.IsAdmin != 1 {
		return c.String(http.StatusBadRequest, "isAdmin認証NG")
	}
	return c.String(http.StatusOK, "認証OK")
}
