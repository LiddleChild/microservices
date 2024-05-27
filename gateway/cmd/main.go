package main

import (
	"fmt"

	"github.com/LiddleChild/microservices/gateway/config"
	userService "github.com/LiddleChild/microservices/gateway/internal/protobuf/user"
	"github.com/LiddleChild/microservices/gateway/internal/user"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.NewConfig()
	fmt.Printf("Running config\n%v\n", cfg)

	conn, err := grpc.NewClient(cfg.UserServiceHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	r := fiber.New()

	userConn := userService.NewUserServiceClient(conn)
	userHandler := user.NewHandler(userConn)
	r.Get("/users", userHandler.GetAllUsers)
	r.Post("/users", userHandler.CreateUser)

	err = r.Listen(fmt.Sprintf(":%v", cfg.Port))
	if err != nil {
		panic(err)
	}
}
