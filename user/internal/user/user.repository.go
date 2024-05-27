package user

import "errors"

type Repository struct {
	users []User
}

func NewRepository() *Repository {
	return &Repository{
		users: make([]User, 0),
	}
}

func (r *Repository) GetAllUsers() []User {
	return r.users
}

func (r *Repository) CreateUser(u User) error {
	for _, user := range r.users {
		if user.Email == u.Email {
			return errors.New("email already exists")
		}
	}

	r.users = append(r.users, u)

	return nil
}
