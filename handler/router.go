package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shinya0226/kensyu/infra/mysql"
	"github.com/shinya0226/kensyu/usecase"
)

func InitRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	sqlHandler := mysql.NewSqlHandler() //sqlHandler
	userRepository := mysql.NewUserRepository(sqlHandler.Conn)
	userUsecase := usecase.NewUserUsecase(userRepository)

	handler := NewUserHandler(userUsecase)
	// v1/users
	e.GET("/", Hello)
	// v1/students/{student_id}
	e.POST("/login", handler.FindSingleRow())

	//e.Logger.Fatal(e.Start(":8080"))

	return e

}
