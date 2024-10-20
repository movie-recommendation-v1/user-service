package storage

import (
	"context"
	pb "github.com/movie-recommendation-v1/user-service/genproto/userservice"
)

type Storage interface {
	User() UserStorageI
	Admin() AdminStorageI
}

type AdminStorageI interface {
	CreateAdmin(ctx context.Context, req *pb.CreateAdminReq) (*pb.CreateAdminRes, error)
	UpdateAdmin(ctx context.Context, req *pb.UpdateAdminReq) (*pb.UpdateAdminRes, error)
	GetAdmin(ctx context.Context, req *pb.GetAdminReq) (*pb.GetAdminRes, error)
	ForgetPassword(ctx context.Context, req *pb.ForgetPasswordReq) (*pb.ForgetPasswordRes, error)
	GetAllAdmins(ctx context.Context, req *pb.GetAllAdminReq) (*pb.GetAllAdminRes, error)
	DeleteAdmin(ctx context.Context, req *pb.DeleteAdminReq) (*pb.DeleteAdminRes, error)
}

type UserStorageI interface {
	Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error)
	RegisterUser(ctx context.Context, req *pb.RegisterUserReq) (*pb.RegisterUserRes, error)
	VerifyUser(ctx context.Context, req *pb.VerifyUserReq) (*pb.VerifyUserRes, error)
	ForgotPassword(ctx context.Context, req *pb.ForgotPasswordReq) (*pb.ForgotPasswordRes, error)
	GetUserByID(ctx context.Context, req *pb.GetUserByIDReq) (*pb.GetUserByIDRes, error)
	GetAllUsers(ctx context.Context, req *pb.GetAllUserReq) (*pb.GetAllUserRes, error)
	UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserRes, error)
}
