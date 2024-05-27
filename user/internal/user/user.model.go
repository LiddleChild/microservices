package user

import "github.com/LiddleChild/microservices/user/internal/protobuf/user"

type User struct {
	Username string
	Email    string
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
