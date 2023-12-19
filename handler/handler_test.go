package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/shinya0226/kensyu/entity"
	"github.com/shinya0226/kensyu/handler"
	"github.com/shinya0226/kensyu/usecase"
)

func TestLoginWithUsecase(t *testing.T) {
	testCase := []struct {
		Description string
		Entity      string //　入力
		Want        string //　出力
		WantErr     bool   //　エラーが出るときはtrue
	}{
		{
			Description: "EmailとPasswordが両方合致",
			Entity:      `{"email":"shinya.yamamoto6@persol-pt.co.jp","password":"yamamo10","name":"山本真也","isAdmin":0}`,
			Want:        `{"email":"shinya.yamamoto6@persol-pt.co.jp","name":"山本真也","isAdmin":0,"access_token":"Anything"}`,
			WantErr:     false,
		},
		{
			Description: "Emailエラーによる不合致",
			Entity:      `{"email":"Emailは違うよ","password":"yamamo10","name":"山本真也","isAdmin":0}`,
			Want:        `{"email":"shinya.yamamoto6@persol-pt.co.jp","name":"山本真也","isAdmin":0,"access_token":"Anything"}`,
			WantErr:     true,
		},
		{
			Description: "Passwordエラーによる不合致",
			Entity:      `{"email":"shinya.yamamoto6@persol-pt.co.jp","password":"Passwordは違うよ","name":"山本真也","isAdmin":0}`,
			Want:        `{"email":"shinya.yamamoto6@persol-pt.co.jp","name":"山本真也","isAdmin":0,"access_token":"Anything"}`,
			WantErr:     true,
		},
		{
			Description: "Nothingエラーによる不合致",
			Entity:      `{"email":"","password":"","name":"","isAdmin":}`,
			Want:        `{"email":"shinya.yamamoto6@persol-pt.co.jp","name":"山本真也","isAdmin":0,"access_token":"Anything"}`,
			WantErr:     true,
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
			ctrl := gomock.NewController(t)
			//　mockの生成
			testMock := handler.NewMockILoginUsecase(ctrl)
			testMock.EXPECT().Login(userEntity).Return(userResponse, nil).AnyTimes()

			req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(tt.Entity))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := handler.LoginWithUsecase(testMock, c)
			if (err != nil) != tt.WantErr {
				t.Errorf("LoginWithUsecase() error = %v, wantErr %v", err, tt.WantErr)
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			// assert.Equal(t, tt.Want+"\n", rec.Body.String())
		})
	}
}
