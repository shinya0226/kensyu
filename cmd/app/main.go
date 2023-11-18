package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/shinya0226/kensyu/handler"
)

func main() {
	fmt.Println("sever start")
	e := echo.New()
	handler.InitRouter()
	e.Logger.Fatal(e.Start(":8080"))
}
