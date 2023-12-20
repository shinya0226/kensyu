package handler_test

import (
	"encoding/json"
	"errors"
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

func TestLoginWithUsecase(t *testing.T) {
	testCase := []struct {
		Description string
		Usecase     func(testMock *handler.MockILoginUsecase)
		Entity      entity.User         //　入力
		Want        usecase.LoginFormat //　出力
		WantErr     bool                //　エラーが出るときはtrue
		WantCode    int
	}{
		{
			Description: "EmailとPasswordが両方合致",
			Entity:      entity.User{Email: "shinya.yamamoto6@persol-pt.co.jp", Password: "yamamo10", Name: "", IsAdmin: 0},
			Want: usecase.LoginFormat{Email: "shinya.yamamoto6@persol-pt.co.jp", Name: "山本真也", IsAdmin: 0,
				AccessToken: "Anything"},
			Usecase: func(testMock *handler.MockILoginUsecase) {
				testMock.EXPECT().Login(entity.User{Email: "shinya.yamamoto6@persol-pt.co.jp", Password: "yamamo10", Name: "山本真也", IsAdmin: 0}).
					Return(usecase.LoginFormat{Email: "shinya.yamamoto6@persol-pt.co.jp", Name: "山本真也", IsAdmin: 0, AccessToken: "Anything"}, nil)
			},
			WantErr:  false,
			WantCode: http.StatusOK,
		},
		{
			Description: "Nothingエラーによる不合致",
			Entity:      entity.User{Email: "", Password: "", Name: "", IsAdmin: 0},
			Want:        usecase.LoginFormat{Email: "", Name: "", IsAdmin: 0, AccessToken: ""},
			// Want:usecase.LoginFormat{Email: "",Name: "",IsAdmin: 0,AccessToken: "Anything"},
			Usecase: func(testMock *handler.MockILoginUsecase) {
				testMock.EXPECT().Login(entity.User{Email: "", Password: "", Name: "", IsAdmin: 0}).
					Return(usecase.LoginFormat{Email: "", Name: "", IsAdmin: 0, AccessToken: ""}, errors.New("empty"))
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
			tt.Usecase(testMock)
			var entity entity.User
			v, err := json.Marshal(entity)
			if err != nil {
				log.Fatal(err)
			}
			req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(v)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			//　Login test
			h := handler.Login(testMock)
			err = h(c)
			assert.Equal(t, tt.WantErr, err != nil)
			assert.Equal(t, tt.WantCode, rec.Code)
		})
	}
}
