package usecase_test

import (
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shinya0226/kensyu/entity"
	. "github.com/shinya0226/kensyu/infra/mysql"
	"github.com/shinya0226/kensyu/usecase"
	"github.com/stretchr/testify/assert"
	"gopkg.in/testfixtures.v1"
)

// login_testの実行
// EmailとPasswordの合致確認
func TestLogin(t *testing.T) {
	type user struct {
		Email    string `json:"Email"`
		Password string `json:"Password"`
		Name     string `json:"Name"`
		IsAdmin  int    `json:"IsAdmin"`
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
			Want:        LoginFormat{"", "", 0, ""},
			WantErr:     true,
		},
		{
			Description: "Nothingエラーによる不合致",
			Entity:      user{"", "", "山本真也", 0},
			Want:        LoginFormat{"", "", 0, ""},
			WantErr:     true,
		},
	}
	//　DB接続
	for _, tt := range testCase {
		t.Run(tt.Description, func(t *testing.T) {
			db, err := ConnectionDB()
			if err != nil {
				log.Fatal(err)
			}
			//　fixtureの設定
			err = testfixtures.LoadFixtures("../testdata/fixtures", db, &testfixtures.MySQLHelper{})
			if err != nil {
				log.Fatal(err)
			}
			userRepo := NewUserRepository(db)
			loginUsecase := usecase.NewLoginUsecase(userRepo)
			got, err := loginUsecase.Login(entity.User(tt.Entity))

			//　errがあるか判別（あるときはtrue,ないときはfalse）
			if (err != nil) != tt.WantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.WantErr)
			}
			//　gotとtt.Wantの中身を比較
			assert.Equal(t, tt.Want.Email, got.Email)
			assert.Equal(t, tt.Want.Name, got.Name)
			assert.Equal(t, tt.Want.IsAdmin, got.IsAdmin)
		})
	}
}
