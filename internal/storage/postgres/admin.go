package postgres

import (
	"context"
	"database/sql"
	"errors"
	pb "github.com/movie-recommendation-v1/user-service/genproto/userservice"
	logger "github.com/movie-recommendation-v1/user-service/internal/log"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

type AdminStorage interface {
	CreateAdmin(ctx context.Context, req *pb.CreateAdminReq) (*pb.CreateAdminRes, error)
	UpdateAdmin(ctx context.Context, req *pb.UpdateAdminReq) (*pb.UpdateAdminRes, error)
	GetAdmin(ctx context.Context, req *pb.GetAdminReq) (*pb.GetAdminRes, error)
	ForgetPassword(ctx context.Context, req *pb.ForgetPasswordReq) (*pb.ForgetPasswordRes, error)
	GetAllAdmins(ctx context.Context, req *pb.GetAllAdminReq) (*pb.GetAllAdminRes, error)
	DeleteAdmin(ctx context.Context, req *pb.DeleteAdminReq) (*pb.DeleteAdminRes, error)
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
	query := `INSERT INTO admins (name, email, password) VALUES ($1, $2, $3)`
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	password, err := hashPassword(req.AdminPassword)
	if err != nil {
		logs.Error("Error hashing password", zap.Error(err))
	}
	_, err = s.db.ExecContext(ctx, query, req.AdminName, req.AdminEmail, password)
	if err != nil {
		logs.Error("Error creating admin", zap.Error(err))
	}
	logs.Info("Successfully created admin")
	return &pb.CreateAdminRes{Message: "Successfully create admin"}, nil
}

func (s *AdminStorageImpl) UpdateAdmin(ctx context.Context, req *pb.UpdateAdminReq) (*pb.UpdateAdminRes, error) {
	query := "UPDATE users SET"
	var args []interface{}
	var updates []string
	argCounter := 1
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	// Check for fields and build query
	if req.AdminReq.Name != "string" && req.AdminReq.Name != "" {
		updates = append(updates, " name = $"+strconv.Itoa(argCounter))
		args = append(args, req.AdminReq.Name)
		argCounter++
	}

	if req.AdminReq.Lastname != "string" && req.AdminReq.Lastname != "" {
		updates = append(updates, " lastname = $"+strconv.Itoa(argCounter))
		args = append(args, req.AdminReq.Lastname)
		argCounter++
	}

	if req.AdminReq.Email != "string" && req.AdminReq.Email != "" {
		updates = append(updates, " email = $"+strconv.Itoa(argCounter))
		args = append(args, req.AdminReq.Email)
		argCounter++
	}

	if len(updates) == 0 {
		return nil, errors.New("no fields to update")
	}

	// Join updates to form the final query
	query += " " + strings.Join(updates, ", ") + " WHERE id = $" + strconv.Itoa(argCounter)
	args = append(args, req.AdminReq.Id)

	// Execute the query
	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		// Log the error and return
		logs.Error("Error updating user: ", zap.Error(err))
		return nil, err
	}

	// Return success response
	return &pb.UpdateAdminRes{AdminRes: req.AdminReq}, nil

}

func (s *AdminStorageImpl) GetAdmin(ctx context.Context, req *pb.GetAdminReq) (*pb.GetAdminRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	query := `select id, name, lastname, email, created_at, updated_at from users where id = $1`
	admin := pb.AdminModel{}
	err = s.db.QueryRowContext(ctx, query, req.AdminId).Scan(&admin.Id, &admin.Name, &admin.Lastname, &admin.Email, &admin.CreatedAt, &admin.UpdatedAt)
	if err != nil {
		logs.Error("Error getting user", zap.Error(err))
		return nil, err
	}
	logs.Info("Successfully got admin")
	resp := pb.GetAdminRes{
		AdminRes: &admin,
	}
	return &resp, nil
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
