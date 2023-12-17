package main_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shinya0226/kensyu/entity"
	"github.com/shinya0226/kensyu/handler"
	"github.com/shinya0226/kensyu/usecase"
)

var userEntity = entity.User{
	Email:    "shinya.yamamoto6@persol-pt.co.jp",
	Password: "yamamo10",
	Name:     "山本真也",
	IsAdmin:  0}

// Login()の出力
var userResponse = usecase.LoginFormat{
	Email:       "shinya.yamamoto6@persol-pt.co.jp",
	Name:        "山本真也",
	IsAdmin:     0,
	AccessToken: "Anything"}

func main_test(t *testing.T) {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// db := mysql.ConnectionDB()
	// userRepo := mysql.NewUserRepository(db)
	// loginUsecase := usecase.NewLoginUsecase(userRepo)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	//　mockの生成
	testMock := handler.NewMockILoginUsecase(ctrl)
	testMock.EXPECT().Login(userEntity).Return(userResponse, nil)

	// ルートを設定
	e.POST("/login", handler.Login(testMock))

	// サーバーをポート番号8080で起動
	e.Logger.Fatal(e.Start(":8080"))
}
