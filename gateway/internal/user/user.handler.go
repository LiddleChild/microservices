package user

import (
	"github.com/LiddleChild/microservices/gateway/internal/protobuf/user"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service user.UserServiceClient
}

func NewHandler(service user.UserServiceClient) *Handler {
	return &Handler{
		service,
	}
}

func (h *Handler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers(c.Context(), nil)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(users.Users)
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	u := &User{}
	err := c.BodyParser(u)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	_, err = h.service.CreateUser(c.Context(), &user.CreateUserRequest{
		User: u.ToServiceModel(),
	})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(201)
}
