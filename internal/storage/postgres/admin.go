package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-redis/redis"
	"strconv"
	"strings"
	"time"

	pb "github.com/movie-recommendation-v1/user-service/genproto/userservice"
	logger "github.com/movie-recommendation-v1/user-service/internal/logger"
	"go.uber.org/zap"
)

type AdminRepo struct {
	db *sql.DB
	rd *redis.Client
}

func NewAdminRepo(db *sql.DB, rd *redis.Client) *AdminRepo {
	return &AdminRepo{
		db: db,
		rd: rd,
	}
}

func (s *AdminRepo) CreateAdmin(ctx context.Context, req *pb.CreateAdminReq) (*pb.CreateAdminRes, error) {
	query := `INSERT INTO admins (name, email, password, role) VALUES ($1, $2, $3, $4)`
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	password, err := hashPassword(req.AdminPassword)
	if err != nil {
		logs.Error("Error hashing password", zap.Error(err))
	}
	_, err = s.db.ExecContext(ctx, query, req.AdminName, req.AdminEmail, password, "admin")
	if err != nil {
		logs.Error("Error creating admin", zap.Error(err))
	}
	logs.Info("Successfully created admin")
	return &pb.CreateAdminRes{Message: "Successfully create admin"}, nil
}

func (s *AdminRepo) UpdateAdmin(ctx context.Context, req *pb.UpdateAdminReq) (*pb.UpdateAdminRes, error) {
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

	//if req.AdminReq.Lastname != "string" && req.AdminReq.Lastname != "" {
	//	updates = append(updates, " lastname = $"+strconv.Itoa(argCounter))
	//	args = append(args, req.AdminReq.Lastname)
	//	argCounter++
	//}

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

func (s *AdminRepo) GetAdmin(ctx context.Context, req *pb.GetAdminReq) (*pb.GetAdminRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	query := `select id, name, lastname, email, created_at, updated_at from users where id = $1`
	admin := pb.AdminModel{}
	err = s.db.QueryRowContext(ctx, query, req.AdminId).Scan(&admin.Id, &admin.Name, &admin.Email, &admin.CreatedAt, &admin.UpdatedAt, &admin.Role)
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

func (s *AdminRepo) ForgetPassword(ctx context.Context, req *pb.ForgetPasswordReq) (*pb.ForgetPasswordRes, error) {
	return nil, nil
}

func (s *AdminRepo) GetAllAdmins(ctx context.Context, req *pb.GetAllAdminReq) (*pb.GetAllAdminRes, error) {
	query := "SELECT id, name, lastname, email, O_CHAR(created_at, 'DD-MM-YYYY') AS created_at, updated_at FROM users WHERE deleted_at = 0"

	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	args := []interface{}{}
	argCounter := 1

	if req.AdminReq.Id != "string" && req.AdminReq.Id != "" {
		query += " AND id = $" + strconv.Itoa(argCounter)
		args = append(args, req.AdminReq.Id)
		argCounter++
	}

	if req.AdminReq.Email != "string" && req.AdminReq.Email != "" {
		query += " AND email = $" + strconv.Itoa(argCounter)
		args = append(args, req.AdminReq.Email)
		argCounter++
	}
	//
	//if req.AdminReq.Lastname != "string" && req.AdminReq.Lastname != "" {
	//	query += " AND lastname = $" + strconv.Itoa(argCounter)
	//	args = append(args, req.AdminReq.Lastname)
	//	argCounter++
	//}

	if req.AdminReq.Name != "string" && req.AdminReq.Name != "" {
		query += " AND name = $" + strconv.Itoa(argCounter)
		args = append(args, req.AdminReq.Name)
		argCounter++
	}

	if req.AdminReq.CreatedAt != "string" && req.AdminReq.CreatedAt != "" {
		t1, err := time.Parse("01-02-2006", req.AdminReq.CreatedAt)
		if err != nil {
			logs.Error("Error parsing date")
			return nil, err
		}
		createdAtInSeconds := t1.Unix()
		query += " AND EXTRACT(EPOCH FROM created_at) > $" + strconv.Itoa(argCounter)

		args = append(args, createdAtInSeconds)
		argCounter++
	}

	// Query execution
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		logs.Error("Error with get all users query")
		return nil, err
	}
	var admins []*pb.AdminModel
	for rows.Next() {
		admin := pb.AdminModel{}

		err = rows.Scan(&admin.Id, &admin.Name, &admin.Email, &admin.CreatedAt, &admin.UpdatedAt, &admin.Role)
		if err != nil {
			logs.Error("Error with get all users")
		}
		admins = append(admins, &admin)
	}

	resp := pb.GetAllAdminRes{
		AdminRes: admins,
	}

	logs.Info("Successfully get all admins")
	return &resp, nil
}

func (s *AdminRepo) DeleteAdmin(ctx context.Context, req *pb.DeleteAdminReq) (*pb.DeleteAdminRes, error) {
	query := "update users set deleted_at = $1 where id = $2"
	t := time.Now().Unix()
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	_, err = s.db.ExecContext(ctx, query, t, req.AdminId)
	if err != nil {
		logs.Error("Error while deleting admin")
	}
	logs.Info("Successfully deleted admin")
	return &pb.DeleteAdminRes{Message: "Success"}, nil

}
