package handler_test

import (
	"errors"
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

	testCase := []struct {
		Description string
		Usecase     func(testMock *handler.MockILoginUsecase)
		Entity      string //　入力
		Want        string //　出力
		WantErr     bool   //　エラーが出るときはtrue
		WantCode    int
	}{
		{
			Description: "EmailとPasswordが両方合致",
			Entity:      `{"email":"shinya.yamamoto6@persol-pt.co.jp","password":"yamamo10","name":"山本真也","isAdmin":0}`,
			Want:        `{"email":"shinya.yamamoto6@persol-pt.co.jp","name":"山本真也","isAdmin":0,"access_token":"Anything"}`,
			Usecase: func(testMock *handler.MockILoginUsecase) {
				testMock.EXPECT().Login(userEntity).Return(userResponse, nil)
			},
			WantErr:  false,
			WantCode: http.StatusOK,
		},
		{
			Description: "Nothingエラーによる不合致",
			Entity:      `{"email":"","password":"","name":"","isAdmin":}`,
			Want:        `{"email":"shinya.yamamoto6@persol-pt.co.jp","name":"山本真也","isAdmin":0,"access_token":"Anything"}`,
			Usecase: func(testMock *handler.MockILoginUsecase) {
				testMock.EXPECT().Login(userEntity).Return(nil, errors.New("Nothingエラー"))
			},
			WantErr:  true,
			WantCode: http.StatusBadRequest,
		},
	}

	for _, tt := range testCase {
		t.Run(tt.Description, func(t *testing.T) {
			e := echo.New()
			ctrl := gomock.NewController(t)
			ctrl.Finish()
			//　mockの生成
			testMock := handler.NewMockILoginUsecase(ctrl)
			//　Usecase test
			tt.Usecase(testMock)

			req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(tt.Entity))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			//　Login test
			h := handler.Login(testMock)
			err := h(c)
			assert.Equal(t, err != nil, tt.WantErr)
			assert.Equal(t, rec.Code, tt.WantErr)
		})
	}
}
