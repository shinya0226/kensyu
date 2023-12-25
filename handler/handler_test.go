package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/shinya0226/kensyu/entity"
	"github.com/shinya0226/kensyu/handler"
	"github.com/shinya0226/kensyu/usecase"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestLogin(t *testing.T) {
	email := "shinya.yamamoto6@persol-pt.co.jp"
	pass := "yamamo10"
	name := "山本真也"
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
			Entity:      entity.User{Email: email, Password: pass, Name: name, IsAdmin: 0},
			Want: usecase.LoginFormat{Email: email, Name: name, IsAdmin: 0,
				AccessToken: "Anything"},
			Usecase: func(testMock *handler.MockILoginUsecase) {
				testMock.EXPECT().Login(entity.User{Email: email, Password: pass, Name: name, IsAdmin: 0}).
					Return(usecase.LoginFormat{Email: email, Name: name, IsAdmin: 0, AccessToken: "Anything"}, nil)
			},
			WantErr:  false,
			WantCode: http.StatusOK,
		},
		{
			Description: "Emailエラーによる不合致",
			Entity:      entity.User{Email: "Emailは違うよ", Password: pass, Name: name, IsAdmin: 0},
			Want:        usecase.LoginFormat{},
			Usecase: func(testMock *handler.MockILoginUsecase) {
				testMock.EXPECT().Login(entity.User{Email: "Emailは違うよ", Password: pass, Name: name, IsAdmin: 0}).
					Return(usecase.LoginFormat{Email: "", Name: "", IsAdmin: 0, AccessToken: ""}, errors.New("Email error"))
			},
			WantErr:  true,
			WantCode: http.StatusOK,
		},
		{
			Description: "Passwordエラーによる不合致",
			Entity:      entity.User{Email: email, Password: "Passwordは違うよ", Name: name, IsAdmin: 0},
			Want:        usecase.LoginFormat{},
			Usecase: func(testMock *handler.MockILoginUsecase) {
				testMock.EXPECT().Login(entity.User{Email: email, Password: "Passwordは違うよ", Name: name, IsAdmin: 0}).
					Return(usecase.LoginFormat{Email: "", Name: "", IsAdmin: 0, AccessToken: ""}, errors.New("Password error"))
			},
			WantErr:  true,
			WantCode: http.StatusOK,
		},
		{
			Description: "Nothingエラーによる不合致",
			Entity:      entity.User{Email: "", Password: "", Name: "", IsAdmin: 0},
			Want:        usecase.LoginFormat{},
			Usecase: func(testMock *handler.MockILoginUsecase) {
				testMock.EXPECT().Login(entity.User{Email: "", Password: "", Name: "", IsAdmin: 0}).
					Return(usecase.LoginFormat{Email: "", Name: "", IsAdmin: 0, AccessToken: ""}, errors.New("empty"))
			},
			WantErr:  false,
			WantCode: http.StatusNotFound,
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
			v, err := json.Marshal(tt.Entity)
			if err != nil {
				log.Fatal(err)
			}
			req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(v))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			//　Login test
			h := handler.Login(testMock)
			err = h(c)
			if (err != nil) != tt.WantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.WantErr)
			}
			assert.Equal(t, tt.WantCode, rec.Code)
		})
	}
}

func TestRestricted(t *testing.T) {
	email := "shinya.yamamoto6@persol-pt.co.jp"
	pass := "yamamo10"
	name := "山本真也"
	testCase := []struct {
		Description string
		Entity      entity.User         //　入力
		Want        usecase.LoginFormat //　出力
		WantErr     bool                //　エラーが出るときはtrue
		WantCode    int
	}{
		{
			Description: "EmailとPasswordが両方合致",
			Entity:      entity.User{Email: email, Password: pass, Name: name, IsAdmin: 0},
			Want: usecase.LoginFormat{Email: email, Name: name, IsAdmin: 0,
				AccessToken: "Anything"},
			WantErr:  false,
			WantCode: http.StatusOK,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.Description, func(t *testing.T) {
			e := echo.New()
			v, err := json.Marshal(tt.Entity)
			if err != nil {
				log.Fatal(err)
			}
			req := httptest.NewRequest(http.MethodGet, "/restricted", bytes.NewReader(v))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.Set("user", "jififijfijf")
			err = handler.Restricted(c)
			if status := rec.Code; status != http.StatusFound {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusFound)
			}
		})
	}
}
