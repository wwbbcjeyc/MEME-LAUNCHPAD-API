// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"context"
	"fmt"
	"meme-launchpad-api/internal/config"
	"meme-launchpad-api/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)
// 服务上下文
type ServiceContext struct {
	Config config.Config

	DB *pgxpool.Pool

	Redis *redis.Client

	// Models
	UserModel *model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 初始化数据库连接
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
		c.Database.SSLMode,
	)

	poolConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse database config: %v", err))
	}

	poolConfig.MaxConns = int32(c.Database.MaxOpenConns)
	poolConfig.MinConns = int32(c.Database.MaxIdleConns)

	db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// 初始化 Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: c.Redis.Pass,
		DB:       c.Redis.DB,
	})

	// 测试 Redis 连接
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		fmt.Printf("Warning: Failed to connect to Redis: %v\n", err)
	}

	return &ServiceContext{
		Config:    c,
		DB:        db,
		Redis:     rdb,
		UserModel: model.NewUserModel(db),
	}
}
