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
