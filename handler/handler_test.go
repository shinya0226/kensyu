package handler_test

import (
	"testing"

	"github.com/shinya0226/kensyu/entity"
	"github.com/shinya0226/kensyu/handler"
	"github.com/stretchr/testify/assert"

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
		Description string      `json:"Description"`
		Entity      user        `json:"Email"` //　入力
		Want        LoginFormat //　出力
		WantErr     bool        //　エラーが出るときはtrue
	}{
		{
			Description: "EmailとPasswordが両方合致",
			Entity:      user{"shinya.yamamoto6@persol-pt.co.jp", "yamamo10", "山本真也", 0},
			Want:        LoginFormat{"shinya.yamamoto6@persol-pt.co.jp", "山本真也", 0, "Anything"},
			WantErr:     false,
		},
		{
			Description: "Emailエラーによる不合致",
			Entity:      user{"Emailは違うよ", "yamamo10", "山本真也", 0},
			Want:        LoginFormat{"", "", 0, ""},
			WantErr:     true,
		},
		{
			Description: "Passwordエラーによる不合致",
			Entity:      user{"shinya.yamamoto6@persol-pt.co.jp", "Passwordは違うよ", "山本真也", 0},
			Want:        LoginFormat{"shinya.yamamoto6@persol-pt.co.jp", "", 0, ""},
			WantErr:     true,
		},
		{
			Description: "Nothingエラーによる不合致",
			Entity:      user{"", "", "山本真也", 0},
			Want:        LoginFormat{"", "", 0, ""},
			WantErr:     true,
		},
	}
	// var userEntity = entity.User{
	// 	Email:    "shinya.yamamoto6@persol-pt.co.jp",
	// 	Password: "yamamo10",
	// 	Name:     "山本真也",
	// 	IsAdmin:  0}

	// // Login()の出力
	// var userResponse = usecase.LoginFormat{
	// 	Email:       "shinya.yamamoto6@persol-pt.co.jp",
	// 	Name:        "山本真也",
	// 	IsAdmin:     0,
	// 	AccessToken: "Anything"}

	for _, tt := range testCase {
		t.Run(tt.Description, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			// defer ctrl.Finish()
			//　mockの生成
			testMock := handler.NewMockILoginUsecase(ctrl)
			testMock.EXPECT().Login(tt.Entity).Return(tt.Want, nil)
			// handler.Login(testMock)
			got, err := testMock.Login(entity.User(tt.Entity))
			//　errがあるか判別（あるときはtrue,ないときはfalse）
			if (err != nil) != tt.WantErr {
				t.Errorf("FindSingleRow() error = %v, wantErr %v", err, tt.WantErr)
			}
			//　gotとtt.Wantの中身を比較
			assert.Equal(t, tt.Want.Email, got.Email)
			assert.Equal(t, tt.Want.Name, got.Name)
			assert.Equal(t, tt.Want.IsAdmin, got.IsAdmin)
		})
	}
	// response := handler.Login(testMock)
}
