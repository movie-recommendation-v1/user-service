package postgres

import (
	"context"
	"database/sql"
	pb "github.com/movie-recommendation-v1/user-service/genproto/userservice"
)

type AdminStorage interface {
	
}
type AdminStorageImpl struct {
	db *sql.DB
}

func NewAdminStorage(db *sql.DB) AdminStorage {
	return &AdminStorageImpl{
		db: db,
	}
}

func (s *AdminStorageImpl) CreateAdmin(ctx context.Context, req *pb.CreateAdminReq) (*pb.CreateAdminRes, error) {
	return nil, nil
}

func (s *AdminStorageImpl) UpdateAdmin(ctx context.Context, req *pb.UpdateAdminReq) (*pb.UpdateAdminRes, error) {
	return nil, nil
}

func (s *AdminStorageImpl) GetAdmin(ctx context.Context, req *pb.GetAdminReq) (*pb.GetAdminRes, error) {
	return nil, nil
}

func (s *AdminStorageImpl) ForgetPassword(ctx context.Context, req *pb.ForgetPasswordReq) (*pb.ForgetPasswordRes, error) {
	return nil, nil
}

func (s *AdminStorageImpl) GetAllAdmins(ctx context.Context, req *pb.GetAllAdminReq) (*pb.GetAllAdminRes, error) {
	return nil, nil
}

func (s *AdminStorageImpl) DeleteAdmin(ctx context.Context, req *pb.DeleteAdminReq) (*pb.DeleteAdminRes, error) {
	return nil, nil
}
