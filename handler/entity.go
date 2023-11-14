package handler

type (
	User1 struct {
		Email    string `json:"Email" form:"Email"`
		Password string `json:"Password" form:"Password"`
		Name     string `json:"Name"`
		IsAdmin  int    `json:"IsAdmin"`
	}
)
