package repository

import (
	"Auth_Service/configs"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func InitRedis(cfg *configs.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.RedUrl,
		DB:   0,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal("failed to connect redis: " + err.Error())
	}

	return client
}

func (r *Repository) SetPhoneCode(phoneNumber, code string) error {
	_, err := r.Redis.Get(context.Background(), phoneNumber).Result()
	if err == nil {
		err = r.Redis.Del(context.Background(), phoneNumber).Err()
		if err != nil {
			return err
		}
	}
	return r.Redis.Set(context.Background(), phoneNumber, code, 3*time.Minute).Err()
}

func (r *Repository) CheckCode(phoneNumber, code string) error {
	result, err := r.Redis.Get(context.Background(), phoneNumber).Result()
	if err != nil {
		return err
	}

	if result != code {
		return errors.New("wrong code")
	}

	return nil
}
