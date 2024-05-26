package main

import (
	"fmt"

	"github.com/LiddleChild/microservices/gateway/config"
	"github.com/LiddleChild/microservices/gateway/internal/user"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.NewConfig()

	r := fiber.New()

	userHandler := user.NewHandler()
	r.Get("/users", userHandler.GetAllUsers)

	err := r.Listen(fmt.Sprintf(":%v", cfg.Port))
	if err != nil {
		panic(err)
	}
}
