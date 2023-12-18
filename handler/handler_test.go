package handler_test

import (
	"net/http"
	"net/http/httptest"
	reflect "reflect"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"

	"github.com/shinya0226/kensyu/entity"
	"github.com/shinya0226/kensyu/handler"
	"github.com/shinya0226/kensyu/usecase"
)

func waitMockCall(t *testing.T, ctrl *gomock.Controller, timeout time.Duration) {
	t.Helper()
	called := false
	done := make(chan interface{})

	go func() {
		time.Sleep(timeout)
		close(done)
	}()

	for {
		select {
		case <-done:
			break
		default:
		}

		var c reflect.Value = reflect.ValueOf(ctrl).Elem()
		var cs reflect.Value = c.FieldByName("expectedCalls").Elem()
		i := cs.FieldByName("expected").MapRange()
		called = true
		for i.Next() {
			calls := i.Value()
			if calls.Len() != 0 {
				called = false
			}
		}
		if called {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	if !called {
		t.Log("missed call in time")
	}
}

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
			testMock.EXPECT().Login(userEntity).Return(userResponse, nil)

			// got, err := testMock.Login(entity.User(tt.Entity))
			handler.Login(testMock)
			waitMockCall(t, ctrl, 5*time.Second)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// assert.Equal(t, tt.Want.Email, got.Email)
			// assert.Equal(t, tt.Want.Name, got.Name)
			// assert.Equal(t, tt.Want.IsAdmin, got.IsAdmin)
			// assert.Equal(t, tt.Want.AccessToken, got.AccessToken)
		})
	}
}

type loginUsecase struct {
	repo entity.IUserRepository
}

func NewLoginUsecase(repo entity.IUserRepository) ILoginUsecase {
	return &loginUsecase{repo: repo}
}

type ILoginUsecase interface {
	Login(e entity.User) (LoginFormat, error)
}

func (u *loginUsecase) Login(e entity.User) (LoginFormat, error) {
	//　該当するユーザーを抽出（found）
	found, err := u.repo.FindSingleRow(e.Email)

	//　出力の型を定義
	logfo := LoginFormat{}
	//　Emailの合致確認
	if err != nil {
		return logfo, err
	}
	logfo.Email = found.Email

	//　Passwordの合致確認
	err = usecase.VerifyPassword(found.Password, e.Password)
	if err != nil {
		return logfo, err
	}
	logfo.Name = found.Name
	logfo.IsAdmin = found.IsAdmin

	//　JWTの作成
	jwtMessage, err := usecase.CreateToken(e.Email)
	if err != nil {
		return logfo, err
	}
	//　出力の型を定義
	logfo.AccessToken = jwtMessage

	return logfo, nil
}

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

	var (
		userJSON = `{"email":"shinya.yamamoto6@persol-pt.co.jp","password":"yamamo10", "name":"山本真也","isAdmin":0}`
	)
	e := echo.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	//　mockの生成
	testMock := handler.NewMockILoginUsecase(ctrl)
	testMock.EXPECT().Login(userEntity).Return(userResponse, nil)

	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(userJSON))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler.LoginWithUsecase(testMock, c)
	// if assert.NoError(t, handler.LoginWithUsecase(testMock, c)) {
	// 	assert.Equal(t, http.StatusOK, rec.Code)
	// 	assert.Equal(t, userResponse, rec.Body.String())
	// }
}

// 　見本
func LoginFunc(u usecase.ILoginUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		return handler.LoginWithUsecase(u, c)
	}
}
func LoginWithUsecase(u usecase.ILoginUsecase, c echo.Context) error {
	eu := new(entity.User)
	logfo := LoginFormat{}

	if err := c.Bind(eu); err != nil {
		return err
	}
	//　Loginの出力をmessageに格納（修正）
	message, err := u.Login(*eu)
	if err != nil {
		return err
	}
	//　formatに追加
	logfo.Email = message.Email
	logfo.Name = message.Name
	logfo.IsAdmin = message.IsAdmin
	logfo.AccessToken = message.AccessToken

	return c.JSON(http.StatusOK, logfo) //　structに詰める
}

type LoginFormat struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	IsAdmin     int    `json:"isAdmin"`
	AccessToken string `json:"access_token"`
}
