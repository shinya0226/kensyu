package entity

type User struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Name     string `json:"Name"`
	IsAdmin  int    `json:"IsAdmin"`
}

type IUserRepository interface {
	FindSingleRow(email string) (User, error)
}
