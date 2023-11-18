package entity

import "database/sql"

type User struct {
	Email    string `json:"Email" form:"Email"`
	Password string `json:"Password" form:"Password"`
	Name     string `json:"Name"`
	IsAdmin  int    `json:"IsAdmin"`
}

func UserFromDomainModel(m *User) *User {
	u := &User{
		Email:    m.Email,
		Password: m.Password,
		Name:     m.Name,
		IsAdmin:  m.IsAdmin,
	}
	return u

}

type IUserRepository interface {
	FindSingleRow(db *sql.DB, Email string) (*User, error)
}
