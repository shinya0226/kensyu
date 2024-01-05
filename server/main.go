package main

import (
	"os"
	
	"log"

	_ "github.com/go-sql-driver/mysql"
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

	db, err := mysql.ConnectionDB()
	if err != nil {
		log.Fatal(err)
	}
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
	r.Use(echojwt.WithConfig(config))
	//　JWT認証
	r.GET("", handler.Allowed)                        //　http://localhost:8080/allowed
	r.GET("/accounts/:page", handler.FetchAccounts()) //　http://localhost:8080/allowed/accounts/:page
	e.Logger.Fatal(e.Start(":8080"))
}
