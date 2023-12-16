package handler_test

import (
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/shinya0226/kensyu/entity"
	"github.com/shinya0226/kensyu/handler"
	"github.com/shinya0226/kensyu/usecase"
)

// handler_testの実行
func TestLogin(t *testing.T) {
	// Login()の入力
	type user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		IsAdmin  int    `json:"isAdmin"`
	}
	type LoginFormat struct {
		Email       string `json:"email"`
		Name        string `json:"name"`
		IsAdmin     int    `json:"isAdmin"`
		AccessToken string `json:"access_token"`
	}
	testCase := []struct {
		Description string      `json:"description"`
		Entity      user        `json:"user"` //　入力
		Want        LoginFormat //　出力
		WantErr     bool        //　エラーが出るときはtrue
	}{
		{
			Description: "EmailとPasswordが両方合致",
			Entity:      user{"shinya.yamamoto6@persol-pt.co.jp", "yamamo10", "山本真也", 0},
			Want:        LoginFormat{"shinya.yamamoto6@persol-pt.co.jp", "山本真也", 0, "Anything"},
			WantErr:     false,
		},
	}
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

	for _, tt := range testCase {
		t.Run(tt.Description, func(t *testing.T) {
			e := echo.New()
			// ctx := context.Background()
			e.Use(middleware.Logger())
			e.Use(middleware.Recover())
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			//　mockの生成
			testMock := handler.NewMockILoginUsecase(ctrl)

			testMock.EXPECT().Login(userEntity).Return(userResponse, nil)

			// next := func(c echo.Context) error {
			// 	return handler.LoginWithUsecase(testMock, c)
			// }

			output, err := testMock.Login(userEntity)
			if err != nil {
				log.Fatal(err)
			}
			log.Fatal(output)

			// handler.LoginFunc(testMock)
			// handler.LoginWithUsecase(testMock, c)

			// req := httptest.NewRequest("POST", "/login", strings.NewReader(""))
			// rec := httptest.NewRecorder()
			// c := e.NewContext(req, rec)

			e.POST("/login", handler.Login(testMock))
			e.Logger.Fatal(e.Start(":8080"))

			// next().
			// 	LoginFunc(testMock)
			//  検証
			// log.Fatal(LoginFunc(testMock(next)(c)))
			// handler.LoginFunc()(c)

		})
	}
}
func LoginFunc(u usecase.ILoginUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		return handler.LoginWithUsecase(u, c)
	}
}
