package main

import (
	"fmt"
	"net"

	"github.com/LiddleChild/microservices/user/config"
	userService "github.com/LiddleChild/microservices/user/internal/protobuf/user"
	"github.com/LiddleChild/microservices/user/internal/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.NewConfig()
	fmt.Printf("Running config\n%v\n", cfg)

	db, err := gorm.Open(postgres.Open(cfg.DbUri), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&user.User{})

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	userRepo := user.NewRepository(db)
	userHandler := user.NewHandler(userRepo)
	userService.RegisterUserServiceServer(grpcServer, userHandler)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Port))
	if err != nil {
		panic(err)
	}

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}

}
