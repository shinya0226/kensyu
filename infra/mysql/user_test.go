package mysql_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/shinya0226/kensyu/infra/mysql"
	"github.com/stretchr/testify/assert"
)

func TestFindSingleRow(t *testing.T) {
	type user struct {
		Email    string `json:"Email"`
		Password string `json:"Password"`
		Name     string `json:"Name"`
		IsAdmin  int    `json:"IsAdmin"`
	}
	testCase := []struct {
		Description string `json:"Description"`
		Email       string `json:"Email"`
		Password    string `json:"Password"`
		Want        user
		WantErr     bool //エラーが出るときはtrue
	}{
		{
			Description: "EmailとPasswordが両方合致",
			Email:       "shinya.yamamoto6@persol-pt.co.jp",
			Password:    "yamamo10",
			Want:        user{"shinya.yamamoto6@persol-pt.co.jp", "yamamo10", "山本真也", 0},
			WantErr:     false,
		},
		{
			Description: "Emailエラーによる不合致",
			Email:       "shinya.yamamoto6@persol-pt",
			Password:    "yamamo10",
			Want:        user{"", "", "", 0},
			WantErr:     true,
		},
		{
			Description: "Passwordエラーによる不合致",
			Email:       "shinya.yamamoto6@persol-pt.co.jp",
			Password:    "yamamo",
			Want:        user{"shinya.yamamoto6@persol-pt.co.jp", "yamamo10", "山本真也", 0},
			WantErr:     false,
		},
		{
			Description: "Nothingエラーによる不合致",
			Email:       "",
			Password:    "",
			Want:        user{"", "", "", 0},
			WantErr:     true,
		},
	}

	for _, tt := range testCase {
		t.Run(tt.Description, func(t *testing.T) {
			db := ConnectionDB()
			userRepo := NewUserRepository(db)
			got, err := userRepo.FindSingleRow(tt.Email)

			//errがあるか判別（あるときはtrue,ないときはfalse）
			if (err != nil) != tt.WantErr {
				t.Errorf("FindSingleRow() error = %v, wantErr %v", err, tt.WantErr)
			}
			//gotとtt.Wantの中身を比較
			assert.Equal(t, got.Email, tt.Want.Email)
			assert.Equal(t, got.Password, tt.Want.Password)
			assert.Equal(t, got.Name, tt.Want.Name)
			assert.Equal(t, got.IsAdmin, tt.Want.IsAdmin)
			return
		})
	}
}
