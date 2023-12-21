package mysql_test

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	. "github.com/shinya0226/kensyu/infra/mysql"
	"github.com/stretchr/testify/assert"
	"gopkg.in/testfixtures.v1"
)

// fixtureのファイルパス
const FixturesPath = "../../testdata/fixtures"

// DBの設定
func prepareTestDatabase() {
	//　db := ConnectionDB()
	db, err := sql.Open("mysql", "atsuser:atspass@tcp(localhost:3306)/kensyu_testing?parseTime=true")
	if err != nil {
		return
	}
	err = testfixtures.LoadFixtures(FixturesPath, db, &testfixtures.MySQLHelper{})
	if err != nil {
		log.Fatal(err)
	}
}

// Emailのみの合致確認
func TestFindSingleRow(t *testing.T) {
	const email = "shinya.yamamoto6@persol-pt.co.jp"
	const pass = "$2a$10$t.3jq0H5hhVQBGG1yxj5nOAUtlXp329t7uwZ.7dA0hoZk0V1zVDAS"
	const name = "山本真也"

	type user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		IsAdmin  int    `json:"isAdmin"`
	}
	testCase := []struct {
		Description string `json:"description"`
		Email       string `json:"email"`
		Password    string `json:"password"`
		Want        user
		WantErr     bool //　エラーが出るときはtrue
	}{
		{
			Description: "EmailとPasswordが両方合致",
			Email:       "shinya.yamamoto6@persol-pt.co.jp",
			Password:    "yamamo10",
			Want:        user{email, pass, name, 0},
			WantErr:     false,
		},
		{
			Description: "Emailエラーによる不合致",
			Email:       "Emailは違うよ",
			Password:    "yamamo10",
			Want:        user{"", "", "", 0},
			WantErr:     true,
		},
		{
			Description: "Passwordエラーによる不合致",
			Email:       "shinya.yamamoto6@persol-pt.co.jp",
			Password:    "Passwordは違うよ",
			Want:        user{email, pass, name, 0},
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
			//　fixtureの設定
			prepareTestDatabase()
			userRepo := NewUserRepository(db)
			got, err := userRepo.FindSingleRow(tt.Email)

			//　errがあるか判別（あるときはtrue,ないときはfalse）
			if (err != nil) != tt.WantErr {
				t.Errorf("FindSingleRow() error = %v, wantErr %v", err, tt.WantErr)
			}
			//　gotとtt.Wantの中身を比較
			assert.Equal(t, tt.Want.Email, got.Email)
			assert.Equal(t, tt.Want.Password, got.Password)
			assert.Equal(t, tt.Want.Name, got.Name)
			assert.Equal(t, tt.Want.IsAdmin, got.IsAdmin)
		})
	}
}
