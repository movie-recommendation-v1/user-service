package service

import (
	"context"
	pb "github.com/movie-recommendation-v1/user-service/genproto/userservice"
	logger "github.com/movie-recommendation-v1/user-service/internal/logger"
	"github.com/movie-recommendation-v1/user-service/internal/storage/postgres"
)

type UserService struct {
	storage *postgres.Storage
	pb.UnimplementedUserServiceServer
}

func NewUserService(db *postgres.Storage) *UserService {
	return &UserService{storage: db}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.storage.Useri.Login(ctx, req)
	if err != nil {
		logs.Error("Error while calling Login")
		return nil, err
	}
	logs.Info("Successfully login!")
	return resp, nil
}

func (s *UserService) VerifyUser(ctx context.Context, req *pb.VerifyUserReq) (*pb.VerifyUserRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.storage.Useri.VerifyUser(ctx, req)
	if err != nil {
		logs.Error("Error while calling VerifyUser")
		return nil, err
	}
	logs.Info("Successfully verify!")
	return resp, nil
}

func (s *UserService) RegisterUser(ctx context.Context, req *pb.RegisterUserReq) (*pb.RegisterUserRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.storage.Useri.RegisterUser(ctx, req)
	if err != nil {
		logs.Error("Error while calling RegisterUser")
		return nil, err
	}
	logs.Info("Successfully register the user!")
	return resp, nil
}
func (s *UserService) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordReq) (*pb.ForgotPasswordRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.storage.Useri.ForgotPassword(ctx, req)
	if err != nil {
		logs.Error("Error while calling ForgotPassword")
		return nil, err
	}
	logs.Info("Successfully send the email!")
	return resp, nil
}
func (s *UserService) GetUserByID(ctx context.Context, req *pb.GetUserByIDReq) (*pb.GetUserByIDRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.storage.Useri.GetUserByID(ctx, req)
	if err != nil {
		logs.Error("Error while calling GetUserByID")
		return nil, err
	}
	logs.Info("Successfully get the user")
	return resp, nil
}
func (s *UserService) GetAllUser(ctx context.Context, req *pb.GetAllUserReq) (*pb.GetAllUserRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.storage.Useri.GetAllUsers(ctx, req)
	if err != nil {
		logs.Error("Error while calling GetAllUsers")
		return nil, err
	}
	logs.Info("Successfully Get all users")
	return resp, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.storage.Useri.UpdateUser(ctx, req)
	if err != nil {
		logs.Error("Error while calling UpdateUser")
		return nil, err
	}
	logs.Info("Successfully updated user")
	return resp, nil
}
