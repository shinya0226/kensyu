package entity

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	IsAdmin  int    `json:"isAdmin"`
}

type IUserRepository interface {
	FindSingleRow(email string) (User, error)
}

type AdminFormat struct {
	IsAdmin int `json:"isAdmin"`
}

type LoginFormat struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	IsAdmin     int    `json:"isAdmin"`
	AccessToken string `json:"access_token"`
}

type DeleteFormat struct {
	Email string `json:"email"`
}

type UpdateFormat struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  int    `json:"isAdmin"`
}
