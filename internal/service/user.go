package service

import (
	"context"
	pb "github.com/movie-recommendation-v1/user-service/genproto/userservice"
	logger "github.com/movie-recommendation-v1/user-service/internal/logger"
	"github.com/movie-recommendation-v1/user-service/internal/storage/postgres"
)

type UserService interface {
	Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error)
	RegisterUser(ctx context.Context, req *pb.RegisterUserReq) (*pb.RegisterUserRes, error)
	ForgotPassword(ctx context.Context, req *pb.ForgotPasswordReq) (*pb.ForgotPasswordRes, error)
	GetUserByID(ctx context.Context, req *pb.GetUserByIDReq) (*pb.GetUserByIDRes, error)
	GetAllUsers(ctx context.Context, req *pb.GetAllUserReq) (*pb.GetAllUserRes, error)
	UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserRes, error)
}

type UserServiceImpl struct {
	auth postgres.UsersStorage
	pb.UnimplementedUserServiceServer
}

func NewUserService(auth UserService) *UserServiceImpl {
	return &UserServiceImpl{auth: auth}
}

func (s *UserServiceImpl) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.auth.Login(ctx, req)
	if err != nil {
		logs.Error("Error while calling Login")
		return nil, err
	}
	logs.Info("Successfully login!")
	return resp, nil
}
func (s *UserServiceImpl) RegisterUser(ctx context.Context, req *pb.RegisterUserReq) (*pb.RegisterUserRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.auth.RegisterUser(ctx, req)
	if err != nil {
		logs.Error("Error while calling RegisterUser")
		return nil, err
	}
	logs.Info("Successfully register the user!")
	return resp, nil
}
func (s *UserServiceImpl) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordReq) (*pb.ForgotPasswordRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.auth.ForgotPassword(ctx, req)
	if err != nil {
		logs.Error("Error while calling ForgotPassword")
		return nil, err
	}
	logs.Info("Successfully send the email!")
	return resp, nil
}
func (s *UserServiceImpl) GetUserByID(ctx context.Context, req *pb.GetUserByIDReq) (*pb.GetUserByIDRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.auth.GetUserByID(ctx, req)
	if err != nil {
		logs.Error("Error while calling GetUserByID")
		return nil, err
	}
	logs.Info("Successfully get the user")
	return resp, nil
}
func (s *UserServiceImpl) GetAllUsers(ctx context.Context, req *pb.GetAllUserReq) (*pb.GetAllUserRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.auth.GetAllUsers(ctx, req)
	if err != nil {
		logs.Error("Error while calling GetAllUsers")
		return nil, err
	}
	logs.Info("Successfully Get all users")
	return resp, nil
}
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.auth.UpdateUser(ctx, req)
	if err != nil {
		logs.Error("Error while calling UpdateUser")
		return nil, err
	}
	logs.Info("Successfully updated user")
	return resp, nil
}
