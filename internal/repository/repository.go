package repository

import (
	"Auth_Service/configs"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Repository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func InitRepository(cfg *configs.Config) Repository {
	return Repository{
		DB:    InitDB(cfg),
		Redis: InitRedis(cfg),
	}
}
