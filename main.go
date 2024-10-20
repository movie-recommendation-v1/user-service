package main

import (
	"fmt"
	"github.com/movie-recommendation-v1/user-service/internal/config"
	"net"

	pb "github.com/movie-recommendation-v1/user-service/genproto/userservice"
	logger "github.com/movie-recommendation-v1/user-service/internal/logger"
	"github.com/movie-recommendation-v1/user-service/internal/service"
	"github.com/movie-recommendation-v1/user-service/internal/storage/postgres"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()
	logs, err := logger.NewLogger()
	if err != nil {
		logs.Error("Error while initializing logger")
		return
	}
	db, err := postgres.ConnectPostgres()
	if err != nil {
		logs.Error("Error while initializing postgres connection")
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.AUTHHOST, cfg.AUTHPORT))
	if err != nil {
		logs.Error("Error while initializing listener")
	}

	defer listener.Close()
	logs.Info(fmt.Sprintf("Server start on port: %d", cfg.AUTHPORT))

	userService := service.NewUserService(db)
	adminService := service.NewAdminService(db)

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, userService)
	pb.RegisterAdminServiceServer(s, adminService)
	if err := s.Serve(listener); err != nil {
		logs.Error("Error while initializing server")
	}
}
