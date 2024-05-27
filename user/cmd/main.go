package main

import (
	"fmt"
	"net"

	"github.com/LiddleChild/microservices/user/config"
	userService "github.com/LiddleChild/microservices/user/internal/protobuf/user"
	"github.com/LiddleChild/microservices/user/internal/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config := config.NewConfig()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.Port))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	userRepo := user.NewRepository()
	userHandler := user.NewHandler(userRepo)
	userService.RegisterUserServiceServer(grpcServer, userHandler)

	reflection.Register(grpcServer)

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
