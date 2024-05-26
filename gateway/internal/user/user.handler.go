package user

import "github.com/gofiber/fiber/v2"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetAllUsers(c *fiber.Ctx) error {
	return c.SendString("getting all users")
}
