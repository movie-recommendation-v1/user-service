package user

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	pb "github.com/movie-recommendation-v1/user-service/genproto/userservice"
	logger "github.com/movie-recommendation-v1/user-service/internal/log"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
)

type UsersStorage interface {
	Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error)
	RegisterUser(ctx context.Context, req *pb.RegisterUserReq) (*pb.RegisterUserRes, error)
	ForgotPassword(ctx context.Context, req *pb.ForgotPasswordReq) (*pb.ForgotPasswordRes, error)
	GetUserByID(ctx context.Context, req *pb.GetUserByIDReq) (*pb.GetUserByIDRes, error)
	GetAllUsers(ctx context.Context, req *pb.GetAllUserReq) (*pb.GetAllUserRes, error)
	UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserRes, error)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type userStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) UsersStorage {
	return &userStorage{db: db}
}

func (s *userStorage) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	query := `select password from users where deleted_at = 0`
	password := ""
	err = s.db.QueryRowContext(ctx, query).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			logs.Error("Error getting data from database")
		}
	}
	id := ""
	err = s.db.QueryRowContext(ctx, "select id from users where deleted_at = 0").Scan(&id)
	if err != nil {
		logs.Error("Error getting data from database")
	}

	status := checkPasswordHash(req.Password, cast.ToString(password))

	if !status {
		logs.Error("Incorrect password")
		return nil, err
	}
	userL, err := s.GetUserByID(ctx, &pb.GetUserByIDReq{Userid: id})
	userM := &pb.UserModel{
		Id:        userL.UserRes.Id,
		Name:      userL.UserRes.Name,
		Lastname:  userL.UserRes.Lastname,
		Email:     userL.UserRes.Email,
		CreatedAt: userL.UserRes.CreatedAt,
		UpdatedAt: userL.UserRes.UpdatedAt,
	}

	return &pb.LoginRes{UserRes: userM}, err

}

func (s *userStorage) RegisterUser(ctx context.Context, req *pb.RegisterUserReq) (*pb.RegisterUserRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	query := "insert into users (id,username, email,password) values ($1, $2, $3, $4);"
	hashpass, err := hashPassword(req.Password)
	if err != nil {
		logs.Error("Error with create user")
		return nil, err
	}
	id := uuid.NewString()
	_, err = s.db.Exec(query, id, req.Name, req.Email, hashpass)
	if err != nil {
		logs.Error("Error with create user")
		return nil, err
	}
	userL, err := s.GetUserByID(ctx, &pb.GetUserByIDReq{Userid: id})
	if err != nil {
		logs.Error("Error with create user")
		return nil, err

	}
	userM := &pb.UserModel{
		Id:        userL.UserRes.Id,
		Name:      userL.UserRes.Name,
		Lastname:  userL.UserRes.Lastname,
		Email:     userL.UserRes.Email,
		CreatedAt: userL.UserRes.CreatedAt,
		UpdatedAt: userL.UserRes.UpdatedAt,
	}

	return &pb.RegisterUserRes{UserRes: userM}, nil
}

func (s *userStorage) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordReq) (*pb.ForgotPasswordRes, error) {
	return nil, nil
}

func (s *userStorage) GetUserByID(ctx context.Context, req *pb.GetUserByIDReq) (*pb.GetUserByIDRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	query := `select id,name, lastname,email, created_at, updated_at from users where id = $1 and deleted_at = 0`

	user := pb.UserModel{}

	err = s.db.QueryRowContext(ctx, query, req.Userid).Scan(&user.Id, &user.Name, &user.Lastname, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		logs.Error("Error with get user")
		return nil, err
	}
	return &pb.GetUserByIDRes{UserRes: &user}, err

}

func (s *userStorage) GetAllUsers(ctx context.Context, req *pb.GetAllUserReq) (*pb.GetAllUserRes, error) {
	query := `select`

	if req.UserReq.Id != "string" && req.UserReq.Id != "" {
		query += " id,"
	}

	if req.UserReq.Email != "string" && req.UserReq.Email != "" {
		query += " email,"
	}
	if req.UserReq.Lastname != "string" && req.UserReq.Lastname != "" {
		query += " lastname,"
	}
	if req.UserReq.Name != "string" && req.UserReq.Name != "" {
		query += " name,"
	}

	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		logs.Error("Error with get all users")
	}
	var users []*pb.UserModel
	for rows.Next() {
		user := pb.UserModel{}

		err = rows.Scan(&user.Id, &user.Name, &user.Lastname, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			logs.Error("Error with get all users")
		}
		users = append(users, &user)
	}
	logs.Info("Successfully get all users")
	return nil, nil

}
func (s *userStorage) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {
	//query := "update"
	return nil, nil
}
