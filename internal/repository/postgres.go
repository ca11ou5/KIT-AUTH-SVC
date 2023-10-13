package repository

import (
	"Auth_Service/configs"
	"Auth_Service/internal/entity"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB(cfg *configs.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort, cfg.DbSslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect db: " + err.Error())
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatal("failed to migrate: " + err.Error())
	}

	return db
}
