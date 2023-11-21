package handler_test

import "testing"

func TestPost(t *testing.T) {
	tests := []struct {
		Description string
		Email       string `json:"Email"`
		Password    string `json:"Password"`
		want        bool
	}{
		{
			Description: "EmailとPasswordが両方合致",
			Email:       "shinya.yamamoto6@persol-pt.co.jp",
			Password:    "yamamo10",
			want:        true,
		},
		{
			Description: "Emailエラーによる不合致",
			Email:       "shinya.yamamoto6@persol-pt",
			Password:    "yamamo10",
			want:        false,
		},
		{
			Description: "Passwordエラーによる不合致",
			Email:       "shinya.yamamoto6@persol-pt.co.jp",
			Password:    "yamamo",
			want:        false,
		},
		{
			Description: "Nothingエラーによる不合致",
			Email:       "",
			Password:    "",
			want:        false,
		},
	}

	db := mysql.ConnectionDB()

	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			//FindSingleRowはbool型で返した方がいいんじゃね？
			if got := tt.FindSingleRow(db,tt.Email);got!=
		})
	}
}
