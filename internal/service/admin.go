package service

import (
	"context"

	pb "github.com/movie-recommendation-v1/user-service/genproto/userservice"
	"github.com/movie-recommendation-v1/user-service/internal/logger"
	"github.com/movie-recommendation-v1/user-service/internal/storage/postgres"
)

type AdminService interface {
	CreateAdmin(ctx context.Context, req *pb.CreateAdminReq) (*pb.CreateAdminRes, error)
	UpdateAdmin(ctx context.Context, req *pb.UpdateAdminReq) (*pb.UpdateAdminRes, error)
	GetAdmin(ctx context.Context, req *pb.GetAdminReq) (*pb.GetAdminRes, error)
	ForgetPassword(ctx context.Context, req *pb.ForgetPasswordReq) (*pb.ForgetPasswordRes, error)
	GetAllAdmins(ctx context.Context, req *pb.GetAllAdminReq) (*pb.GetAllAdminRes, error)
	DeleteAdmin(ctx context.Context, req *pb.DeleteAdminReq) (*pb.DeleteAdminRes, error)
}
type adminService struct {
	admin postgres.AdminStorage
	pb.UnimplementedAdminServiceServer
}

func NewAdminStorage(admin postgres.AdminStorage) AdminService {
	return &adminService{
		admin: admin,
	}
}

func (s *adminService) CreateAdmin(ctx context.Context, req *pb.CreateAdminReq) (*pb.CreateAdminRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp , err := s.admin.CreateAdmin(ctx, req)
	if err != nil {
		logs.Error("Error while calling CreateAdmin")
	}
	logs.Info("Successfully create admin")
	return resp, nil
}

func (s *adminService) UpdateAdmin(ctx context.Context, req *pb.UpdateAdminReq) (*pb.UpdateAdminRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	resp, err := s.admin.UpdateAdmin(ctx, req)
	if err != nil {
		logs.Error("Error while calling Update Admin")
	}
	logs.Info("Successfully update admin")
	return resp, nil
}

func (s *adminService) GetAdmin(ctx context.Context, req *pb.GetAdminReq) (*pb.GetAdminRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	
	resp , err := s.admin.GetAdmin(ctx,req)
	if err != nil {
		logs.Error("Error while calling Get Admin")
	}
	logs.Info("Successfully get admin")
	return resp, nil
}

func (s *adminService) ForgetPassword(ctx context.Context, req *pb.ForgetPasswordReq) (*pb.ForgetPasswordRes, error) {
	return nil, nil
}

func (s *adminService) GetAllAdmins(ctx context.Context, req *pb.GetAllAdminReq) (*pb.GetAllAdminRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	resp, err := s.admin.GetAllAdmins(ctx, req)
	if err != nil {
		logs.Error("Error while calling GetAllAdmins")
	}
	logs.Info("Successfully get all admins")
	return resp, nil
}

func (s *adminService) DeleteAdmin(ctx context.Context, req *pb.DeleteAdminReq) (*pb.DeleteAdminRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	resp, err := s.admin.DeleteAdmin(ctx, req)
	if err != nil {
		logs.Error("Error while deleting admin")
	}
	logs.Info("Successfully delete admin")
	return resp,nil
}
