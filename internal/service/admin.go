package service

import (
	"context"

	pb "github.com/movie-recommendation-v1/user-service/genproto/userservice"
	"github.com/movie-recommendation-v1/user-service/internal/logger"
	"github.com/movie-recommendation-v1/user-service/internal/storage/postgres"
)

type AdminService struct {
	storage *postgres.Storage
	pb.UnimplementedAdminServiceServer
}

func NewAdminService(db *postgres.Storage) *AdminService {
	return &AdminService{
		storage: db,
	}
}

func (s *AdminService) CreateAdmin(ctx context.Context, req *pb.CreateAdminReq) (*pb.CreateAdminRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.storage.Admini.CreateAdmin(ctx, req)
	if err != nil {
		logs.Error("Error while calling CreateAdmin")
	}
	logs.Info("Successfully create admin")
	return resp, nil
}

func (s *AdminService) UpdateAdmin(ctx context.Context, req *pb.UpdateAdminReq) (*pb.UpdateAdminRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	resp, err := s.storage.Admini.UpdateAdmin(ctx, req)
	if err != nil {
		logs.Error("Error while calling Update Admin")
	}
	logs.Info("Successfully update admin")
	return resp, nil
}

func (s *AdminService) GetAdmin(ctx context.Context, req *pb.GetAdminReq) (*pb.GetAdminRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	resp, err := s.storage.Admini.GetAdmin(ctx, req)
	if err != nil {
		logs.Error("Error while calling Get Admin")
	}
	logs.Info("Successfully get admin")
	return resp, nil
}

func (s *AdminService) ForgetPassword(ctx context.Context, req *pb.ForgetPasswordReq) (*pb.ForgetPasswordRes, error) {
	return nil, nil
}

func (s *AdminService) GetAllAdmins(ctx context.Context, req *pb.GetAllAdminReq) (*pb.GetAllAdminRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	resp, err := s.storage.Admini.GetAllAdmins(ctx, req)
	if err != nil {
		logs.Error("Error while calling GetAllAdmins")
	}
	logs.Info("Successfully get all admins")
	return resp, nil
}

func (s *AdminService) DeleteAdmin(ctx context.Context, req *pb.DeleteAdminReq) (*pb.DeleteAdminRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	resp, err := s.storage.Admini.DeleteAdmin(ctx, req)
	if err != nil {
		logs.Error("Error while deleting admin")
	}
	logs.Info("Successfully delete admin")
	return resp, nil
}
