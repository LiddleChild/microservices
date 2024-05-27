package user

import "github.com/LiddleChild/microservices/gateway/internal/protobuf/user"

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func FromServiceModel(u *user.User) User {
	return User{
		Username: u.Username,
		Email:    u.Email,
	}
}

func (u User) ToServiceModel() *user.User {
	return &user.User{
		Username: u.Username,
		Email:    u.Email,
	}
}
