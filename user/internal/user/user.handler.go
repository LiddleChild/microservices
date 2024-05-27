package user

import (
	"context"
	"errors"

	"github.com/LiddleChild/microservices/user/internal/protobuf/user"
)

type Handler struct {
	user.UnimplementedUserServiceServer
	repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) GetAllUsers(c context.Context, _ *user.GetAllUsersRequest) (*user.GetAllUsersResponse, error) {
	res := &user.GetAllUsersResponse{
		Users: make([]*user.User, 0),
	}

	var users []User
	err := h.repo.GetAllUsers(&users)
	if err != nil {
		return nil, err
	}

	for _, u := range users {
		res.Users = append(res.Users, u.ToServiceModel())
	}

	return res, nil
}

func (h *Handler) CreateUser(c context.Context, u *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	userModel := FromServiceModel(u.User)

	if len(userModel.Email) == 0 || len(userModel.Username) == 0 {
		return nil, errors.New("user fields cannot be empty")
	}

	err := h.repo.CreateUser(&userModel)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
