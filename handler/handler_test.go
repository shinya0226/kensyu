package handler_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/shinya0226/kensyu/entity"
	"github.com/shinya0226/kensyu/handler"
	"github.com/shinya0226/kensyu/usecase"

	"github.com/golang/mock/gomock"
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
	var userResponse = LoginFormat{
		Email:       "shinya.yamamoto6@persol-pt.co.jp",
		Name:        "山本真也",
		IsAdmin:     0,
		AccessToken: "Anything"}

	for _, tt := range testCase {
		t.Run(tt.Description, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			// defer ctrl.Finish()
			//　mockの生成
			testMock := handler.NewMockILoginUsecase(ctrl)
			testMock.EXPECT().Login(entity.User(userEntity)).Return(usecase.LoginFormat(userResponse), nil)
			// handler.Login(testMock)
			handler.Login(testMock)
		})
	}
	// response := handler.Login(testMock)
}
