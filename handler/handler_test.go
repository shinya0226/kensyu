package handler_test

import (
	"log"
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
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			//　mockの生成
			testMock := handler.NewMockILoginUsecase(ctrl)
			testMock.EXPECT().Login(userEntity).Return(userResponse, nil).AnyTimes()
			// got, err := testMock.Login(entity.User(tt.Entity))
			handler.Login(testMock)
			got, err := testMock.Login(userEntity)
			if err != nil {
				log.Fatal(err)
			}
			// next := func(c echo.Context) error {
			// 	return handler.LoginWithUsecase(testMock, c)
			// }
			assert.Equal(t, tt.Want.Email, got.Email)
			assert.Equal(t, tt.Want.Name, got.Name)
			assert.Equal(t, tt.Want.IsAdmin, got.IsAdmin)
			assert.Equal(t, tt.Want.AccessToken, got.AccessToken)
		})
	}
}

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
			ctrl := gomock.NewController(t)
			//　mockの生成
			testMock := handler.NewMockILoginUsecase(ctrl)
			testMock.EXPECT().Login(userEntity).Return(userResponse, nil).AnyTimes()

			req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(tt.Entity))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if assert.NoError(t, handler.LoginWithUsecase(testMock, c)) {
				assert.Equal(t, http.StatusOK, rec.Code)
				assert.Equal(t, tt.Want+"\n", rec.Body.String())
			}
		})
	}
}
