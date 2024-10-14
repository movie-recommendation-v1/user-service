package service

import (
	"context"
	"database/sql"
	pb "github.com/movie-recommendation-v1/user-service/genproto/userservice"
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
	db *sql.DB
}

func NewAdminStorage(db *sql.DB) AdminService {
	return &adminService{
		db: db,
	}
}

func (s *adminService) CreateAdmin(ctx context.Context, req *pb.CreateAdminReq) (*pb.CreateAdminRes, error) {
	return nil, nil
}

func (s *adminService) UpdateAdmin(ctx context.Context, req *pb.UpdateAdminReq) (*pb.UpdateAdminRes, error) {
	return nil, nil
}

func (s *adminService) GetAdmin(ctx context.Context, req *pb.GetAdminReq) (*pb.GetAdminRes, error) {
	return nil, nil
}

func (s *adminService) ForgetPassword(ctx context.Context, req *pb.ForgetPasswordReq) (*pb.ForgetPasswordRes, error) {
	return nil, nil
}

func (s *adminService) GetAllAdmins(ctx context.Context, req *pb.GetAllAdminReq) (*pb.GetAllAdminRes, error) {
	return nil, nil
}

func (s *adminService) DeleteAdmin(ctx context.Context, req *pb.DeleteAdminReq) (*pb.DeleteAdminRes, error) {
	return nil, nil
}
