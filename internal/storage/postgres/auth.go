package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	pb "github.com/movie-recommendation-v1/user-service/genproto/userservice"
	h "github.com/movie-recommendation-v1/user-service/helper"
	"github.com/movie-recommendation-v1/user-service/internal/config"
	logger "github.com/movie-recommendation-v1/user-service/internal/logger"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	_ "math/rand"
	"strconv"
	"strings"
	"time"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type UserRepo struct {
	db *sql.DB
	rd *redis.Client
}

func NewUserRepo(db *sql.DB, rd *redis.Client) *UserRepo {
	return &UserRepo{
		db: db,
		rd: rd,
	}
}

func (s *UserRepo) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	logs, err := logger.NewLogger()
	query := `
		SELECT
			id,
			name,
			email,
			img_url,
			password,
			role,
			created_at,
			updated_at
		FROM
			users
		WHERE
			email = $1 AND deleted_at = 0
	`
	res := pb.LoginRes{
		UserRes: &pb.UserModel{},
	}
	var password string
	err = s.db.QueryRow(query, req.Email).Scan(
		&res.UserRes.Id,
		&res.UserRes.Name,
		&res.UserRes.Email,
		&password,
		&res.UserRes.Role,
		&res.UserRes.CreatedAt,
		&res.UserRes.UpdatedAt,
	)
	if err != nil {
		logs.Error("Error with scan:", zap.Error(err))
		return nil, err
	}
	ok := checkPasswordHash(req.Password, password)
	if !ok {
		return nil, fmt.Errorf("password not correct")
	}

	return &pb.LoginRes{UserRes: res.UserRes}, err

}

type UserModel struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *UserRepo) RegisterUser(ctx context.Context, req *pb.RegisterUserReq) (*pb.RegisterUserRes, error) {
	logs, err := logger.NewLogger()
	println(req.Name)
	if err != nil {
		return nil, err
	}
	code := rand.Intn(899999) + 100000
	cfg := config.Load()
	email := cfg.EMAIL
	email_key := cfg.EMAILSECREDKEY
	err = h.SendVerificationCode(h.Params{
		From:     email,
		Password: email_key,
		To:       req.Email,
		Message:  fmt.Sprintf("Hi %s, your verification code is %d", req.Name, code),
		Code:     fmt.Sprint(code),
	})
	if err != nil {
		logs.Error("Error with send email wrong email type:", zap.Error(err))
		return nil, err
	}
	user := &UserModel{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	fmt.Println("Name: ", user.Name)
	user1, err := json.Marshal(&user)
	if err != nil {
		logs.Error("Error marshaling user data:", zap.Error(err))
		return nil, err
	}
	err = s.rd.Set(fmt.Sprint(code), user1, time.Minute*5).Err() // Increased expiration time to 5 minutes
	if err != nil {
		logs.Error("Error with redis set:", zap.Error(err))
		return nil, err
	}
	return &pb.RegisterUserRes{Message: "Successfully registered"}, nil
}

func (s *UserRepo) VerifyUser(ctx context.Context, req *pb.VerifyUserReq) (*pb.VerifyUserRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	user1, err := s.rd.Get(req.SmsCode).Result()
	if err != nil {
		if err == redis.Nil {
			logs.Error("Verification code not found or expired")
			return nil, errors.New("verification code not found or expired")
		}
		logs.Error("Error with redis get:", zap.Error(err))
		return nil, err
	}

	var user UserModel
	err = json.Unmarshal([]byte(user1), &user)
	if err != nil {
		logs.Error("Error with redis unmarshal:", zap.Error(err))
		return nil, err
	}

	if user.Name == "" || user.Email == "" || user.Password == "" {
		logs.Error("Invalid user data")
		return nil, errors.New("invalid user data")
	}

	query := "INSERT INTO users (id, name, email, password, img_url) VALUES ($1, $2, $3, $4, $5);"
	hashpass, err := hashPassword(user.Password)
	if err != nil {
		logs.Error("Error hashing password")
		return nil, err
	}

	id := uuid.NewString()
	_, err = s.db.ExecContext(ctx, query, id, user.Name, user.Email, hashpass, "empty")
	if err != nil {
		logs.Error("Error creating user")
		return nil, err
	}

	err = s.rd.Del(req.SmsCode).Err()
	if err != nil {
		logs.Error("Error deleting verification code from Redis")
	}

	userGet, err := s.GetUserByID(ctx, &pb.GetUserByIDReq{Userid: id})
	if err != nil {
		logs.Error("Error getting created user")
		return nil, err
	}

	userM := &pb.UserModel{
		Id:        userGet.UserRes.Id,
		Name:      userGet.UserRes.Name,
		Email:     userGet.UserRes.Email,
		Role:      userGet.UserRes.Role,
		CreatedAt: userGet.UserRes.CreatedAt,
		UpdatedAt: userGet.UserRes.UpdatedAt,
	}

	return &pb.VerifyUserRes{
		Res: userM,
	}, nil
}

//func (s *UserRepo) RegisterUser(ctx context.Context, req *pb.RegisterUserReq) (*pb.RegisterUserRes, error) {
//	logs, err := logger.NewLogger()
//	if err != nil {
//		return nil, err
//	}
//
//	query := "insert into users (id,name, email,password) values ($1, $2, $3, $4);"
//	hashpass, err := hashPassword(req.Password)
//	if err != nil {
//		logs.Error("Error with create user")
//		return nil, err
//	}
//	id := uuid.NewString()
//	_, err = s.db.Exec(query, id, req.Name, req.Email, hashpass)
//	if err != nil {
//		logs.Error("Error with create user")
//		return nil, err
//	}
//	userL, err := s.GetUserByID(ctx, &pb.GetUserByIDReq{Userid: id})
//	if err != nil {
//		logs.Error("Error with create user")
//		return nil, err
//
//	}
//	userM := &pb.UserModel{
//		Id:        userL.UserRes.Id,
//		Name:      userL.UserRes.Name,
//		Email:     userL.UserRes.Email,
//		Role:      userL.UserRes.Role,
//		CreatedAt: userL.UserRes.CreatedAt,
//		UpdatedAt: userL.UserRes.UpdatedAt,
//	}
//
//	return &pb.RegisterUserRes{UserRes: userM}, nil
//}

func (s *UserRepo) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordReq) (*pb.ForgotPasswordRes, error) {
	//logs, err := logger.NewLogger()
	//if err != nil {
	//
	//	return nil, err
	//}
	return nil, nil
}

func (s *UserRepo) GetUserByID(ctx context.Context, req *pb.GetUserByIDReq) (*pb.GetUserByIDRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	query := `select id,name,email, role, created_at, updated_at from users where id = $1 and deleted_at = 0`

	user := pb.UserModel{}

	err = s.db.QueryRowContext(ctx, query, req.Userid).Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Role)
	if err != nil {
		logs.Error("Error with get user")
		return nil, err
	}
	return &pb.GetUserByIDRes{UserRes: &user}, err

}

func (s *UserRepo) GetAllUsers(ctx context.Context, req *pb.GetAllUserReq) (*pb.GetAllUserRes, error) {
	query := "SELECT id, name, email, O_CHAR(created_at, 'DD-MM-YYYY') AS created_at, updated_at FROM users WHERE deleted_at = 0"

	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	args := []interface{}{}
	argCounter := 1

	if req.UserReq.Id != "string" && req.UserReq.Id != "" {
		query += " AND id = $" + strconv.Itoa(argCounter)
		args = append(args, req.UserReq.Id)
		argCounter++
	}

	if req.UserReq.Email != "string" && req.UserReq.Email != "" {
		query += " AND email = $" + strconv.Itoa(argCounter)
		args = append(args, req.UserReq.Email)
		argCounter++
	}

	//if req.UserReq.Lastname != "string" && req.UserReq.Lastname != "" {
	//	query += " AND lastname = $" + strconv.Itoa(argCounter)
	//	args = append(args, req.UserReq.Lastname)
	//	argCounter++
	//}

	if req.UserReq.Name != "string" && req.UserReq.Name != "" {
		query += " AND name = $" + strconv.Itoa(argCounter)
		args = append(args, req.UserReq.Name)
		argCounter++
	}

	if req.UserReq.CreatedAt != "string" && req.UserReq.CreatedAt != "" {
		t1, err := time.Parse("01-02-2006", req.UserReq.CreatedAt)
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
	var users []*pb.UserModel
	for rows.Next() {
		user := pb.UserModel{}

		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.Role, &user.UpdatedAt)
		if err != nil {
			logs.Error("Error with get all users")
		}
		users = append(users, &user)
	}

	resp := pb.GetAllUserRes{
		UserRes: users,
	}

	logs.Info("Successfully get all users")
	return &resp, nil

}
func (s *UserRepo) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {
	query := "UPDATE users SET"
	var args []interface{}
	var updates []string
	argCounter := 1
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	// Check for fields and build query
	if req.UserReq.Name != "string" && req.UserReq.Name != "" {
		updates = append(updates, " name = $"+strconv.Itoa(argCounter))
		args = append(args, req.UserReq.Name)
		argCounter++
	}

	//if req.UserReq.Lastname != "string" && req.UserReq.Lastname != "" {
	//	updates = append(updates, " lastname = $"+strconv.Itoa(argCounter))
	//	args = append(args, req.UserReq.Lastname)
	//	argCounter++
	//}

	if req.UserReq.Email != "string" && req.UserReq.Email != "" {
		updates = append(updates, " email = $"+strconv.Itoa(argCounter))
		args = append(args, req.UserReq.Email)
		argCounter++
	}

	// You can add more fields here...

	// If no fields are provided, return an error
	if len(updates) == 0 {
		return nil, errors.New("no fields to update")
	}

	// Join updates to form the final query
	query += " " + strings.Join(updates, ", ") + " WHERE id = $" + strconv.Itoa(argCounter)
	args = append(args, req.UserReq.Id)

	// Execute the query
	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		// Log the error and return
		logs.Error("Error updating user: ", zap.Error(err))
		return nil, err
	}

	// Return success response
	return &pb.UpdateUserRes{UserRes: req.UserReq}, nil
}
