package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// フィクスチャーを読み込むための型の定義
type Fixture struct {
	TestCases []testCase `json:"test_cases"`
}
type testCase struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	want     bool   `json:"want"`
}

// GetHandlerの実装
func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "お仕事おつかれ様")
}

// PostHandlerの実装
func PostHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("testdata/fixture.json")
	if err != nil {
		log.Fatal(err)
	}
	f := new(Fixture)
	if err := json.Unmarshal(b, f); err != nil {
		log.Fatal(err)
	}

}

// GETメソッドのテスト
func TestGetHandler(t *testing.T) {
	// テスト用のリクエスト作成
	req := httptest.NewRequest("GET", "http://localhost:8080", nil)
	// テスト用のレスポンス作成
	res := httptest.NewRecorder()

	// ハンドラーの実行
	GetHandler(res, req)

	// レスポンスのステータスコードのテスト
	if res.Code != http.StatusOK {
		t.Errorf("invalid code: %d", res.Code)
	} else {
		t.Log("http.StatusOK")
	}

	// レスポンスのボディのテスト
	if res.Body.String() != "お仕事おつかれ様" {
		t.Errorf("invalid response: %#v", res)
	} else {
		t.Log("response body Ok")
	}

}

// POSTメソッドのテスト
func TestPostHandler(t *testing.T) {
	// fixtureの読み込み
	b, err := ioutil.ReadFile("testdata/fixture.json")
	if err != nil {
		log.Fatal(err)
	}
	f := new(Fixture)
	if err := json.Unmarshal(b, f); err != nil {
		log.Fatal(err)
	}

	// テスト用のリクエスト作成
	req := httptest.NewRequest("POST", "http://localhost:8080/login", nil)
	// テスト用のレスポンス作成
	res := httptest.NewRecorder()

	// ハンドラーの実行
	PostHandler(res, req)

	// レスポンスのステータスコードのテスト
	if res.Code != http.StatusOK {
		t.Errorf("invalid code: %d", res.Code)
	} else {
		t.Log("http.StatusOK")
	}

	var buffer bytes.Buffer
	io.WriteString(&buffer, "bytes.Buffer example3")
	fmt.Println(buffer.String())

	//t.Log(res.Body.Bytes())

	// // レスポンスのボディのテスト
	// if res.Body.String() != "helllo" {
	// 	t.Errorf("invalid response: %#v", res)
	// } else {
	// 	t.Log("response body Ok")
	// }

}

func TestLogin(t *testing.T) {
	cases := map[string]struct {
		in   User
		want bool
	}{
		"OK":  {in: User{Email: "shinya.yamamoto6@persol-pt.co.jp", Password: "yamamo10"}, want: true},
		"NO1": {in: User{Email: "shinya.yamamoto6@persol-pt.co", Password: "yamamo10"}, want: false},
		"NO2": {in: User{Email: "shinya.yamamoto6@persol-pt.co.jp", Password: "yamamo"}, want: false},
	}
	for k, tt := range cases {
		t.Run(k, func(t *testing.T) {
			got := tt.in.Login()
			if tt.want != got {
				t.Errorf("want: Email(%s) = %v, got: %v", tt.in.Email, tt.want, got)
				t.Errorf("want: Password(%s) = %v, got: %v", tt.in.Password, tt.want, got)
			}
		})
	}
}

func (u *User) Login() bool {
	if u.Email == "shinya.yamamoto6@persol-pt.co.jp" && u.Password == "yamamo10" {
		return true
	} else {
		return false
	}

}
