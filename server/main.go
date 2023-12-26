package main

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shinya0226/kensyu/handler"
	"github.com/shinya0226/kensyu/infra/mysql"
	"github.com/shinya0226/kensyu/usecase"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := mysql.ConnectionDB()
	userRepo := mysql.NewUserRepository(db)
	loginUsecase := usecase.NewLoginUsecase(userRepo)

	e.POST("/login", handler.Login(loginUsecase))

	r := e.Group("/allowed")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwt.MapClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}
	//　認証
	r.Use(echojwt.WithConfig(config))

	//　JWT認証
	r.GET("", handler.Allowed)                           //　http://localhost:8080/allowed
	r.GET("/accounts/:page", handler.FetchAccounts())    // http://localhost:8080/allowed/accounts/1
	r.POST("/account/new", handler.CreateAccount())      // http://localhost:8080/allowed/account/new
	r.DELETE("/account/delete", handler.DeleteAccount()) // http://localhost:8080/restricted/account/delete
	r.PUT("/account/update", handler.UpdateAccount())    // http://localhost:8080/restricted/account/update

	e.Logger.Fatal(e.Start(":8080"))
}
