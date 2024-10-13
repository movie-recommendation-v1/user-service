package service

import (
	"context"
	pb "github.com/movie-recommendation-v1/user-service/genproto/userservice"
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
	auth UserService
	pb.UnimplementedUserServiceServer
}

func NewUserService(auth UserService) *UserServiceImpl {
	return &UserServiceImpl{auth: auth}
}

func (s *UserServiceImpl) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	return nil, nil
}
func (s *UserServiceImpl) RegisterUser(ctx context.Context, req *pb.RegisterUserReq) (*pb.RegisterUserRes, error) {
	return nil, nil
}
func (s *UserServiceImpl) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordReq) (*pb.ForgotPasswordRes, error) {
	return nil, nil
}
func (s *UserServiceImpl) GetUserByID(ctx context.Context, req *pb.GetUserByIDReq) (*pb.GetUserByIDRes, error) {
	return nil, nil
}
func (s *UserServiceImpl) GetAllUsers(ctx context.Context, req *pb.GetAllUserReq) (*pb.GetAllUserRes, error) {
	return nil, nil
}
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {
	return nil, nil
}
