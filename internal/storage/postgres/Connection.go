package postgres

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
	"github.com/movie-recommendation-v1/user-service/internal/config"
	logger "github.com/movie-recommendation-v1/user-service/internal/logger"
	"github.com/movie-recommendation-v1/user-service/internal/storage"
)

type Storage struct {
	Db     *sql.DB
	Rd     *redis.Client
	Useri  storage.UserStorageI
	Admini storage.AdminStorageI
}

func ConnectPostgres() (*Storage, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	conf := config.Load()
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.DBHOST, conf.DBPORT, conf.DBUSER, conf.DBPASSWORD, conf.DBNAME)
	db, err := sql.Open("postgres", dns)
	if err != nil {
		logs.Error("Error while connecting postgres") //hjhghjghjsdf,gjhsdjkfhgjksdkgsk
	}
	//err = db.Ping()
	//if err != nil {
	//	logs.Error("Errolr while pinging postgres")12345
	//}
	logs.Info("Successfully connected to postgres")

	rd := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", conf.REDISHOST, conf.REDISPORT),
	})
	admin := NewAdminRepo(db, rd)
	user := NewUserRepo(db, rd)
	return &Storage{
		Db:     db,
		Rd:     rd,
		Useri:  user,
		Admini: admin,
	}, nil
}

func (stg *Storage) User() *storage.UserStorageI {
	if stg.Useri == nil {
		stg.Useri = NewUserRepo(stg.Db, stg.Rd)
	}
	return &stg.Useri
}

func (stg *Storage) Admin() *storage.AdminStorageI {
	if stg.Admini == nil {
		stg.Admini = NewAdminRepo(stg.Db, stg.Rd)
	}
	return &stg.Admini
}
