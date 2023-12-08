package main

import (
	"net/http"
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
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := mysql.ConnectionDB()

	userRepo := mysql.NewUserRepository(db)
	loginUsecase := usecase.NewLoginUsecase(userRepo)

	//ログイン処理
	e.POST("/login", handler.Login(loginUsecase))
	// 初期画面
	e.GET("/", handler.Hello)

	//アカウント一覧取得処理
	r := e.Group("/restricted")

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwt.MapClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}
	r.Use(echojwt.WithConfig(config))
	//JWT認証
	r.GET("", restricted)                           //http://localhost:8080/restricted
	r.GET("/accounts/:page", handler.GetAccounts()) // http://localhost:8080/restricted/accounts/1

	// サーバーをポート番号8080で起動
	e.Logger.Fatal(e.Start(":8080"))

}

// JWT認証
func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token) //token
	usecase.VerifyToken(user.Raw)

	return c.String(http.StatusOK, "JWT認証OK")
}
