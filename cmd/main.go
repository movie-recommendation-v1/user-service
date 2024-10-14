package main

import (
	"net"

	pb "github.com/movie-recommendation-v1/user-service/genproto/userservice"
	logger "github.com/movie-recommendation-v1/user-service/internal/logger"
	"github.com/movie-recommendation-v1/user-service/internal/service"
	"github.com/movie-recommendation-v1/user-service/internal/storage/postgres"
	"google.golang.org/grpc"
)
	
func main() {

	logs, err := logger.NewLogger()
	if err != nil {
		logs.Error("Error while initializing logger")
	}

	db, err := postgres.ConnectPostgres()
	if err != nil {
		logs.Error("Error while initializing postgres connection")
	}
	defer db.Close()
	listener, err := net.Listen("tcp", ":5051")
	if err != nil {
		logs.Error("Error while initializing listener")
	}
	defer listener.Close()
	logs.Info("Server start on port 5051")

	userStorage := postgres.NewUserStorage(db)
	userService := service.NewUserService(userStorage)

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, userService)

	if err := s.Serve(listener); err != nil {
		logs.Error("Error while initializing server")
	}
	logs.Info("Bashoi Nahoi")
}
