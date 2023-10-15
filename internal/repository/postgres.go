package repository

import (
	"Auth_Service/configs"
	"Auth_Service/internal/entity"
	"errors"
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

func (r *Repository) CreateUser(user *entity.User) error {
	dbUser := entity.User{}
	result := r.DB.Where("phone_number = ?", user.PhoneNumber).First(&dbUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		res := r.DB.Create(user)
		return res.Error
	}

	if user.Confirmed == false {
		user.Id = dbUser.Id
		res := r.DB.Save(user)
		return res.Error
	}

	return errors.New("the phone number is already taken")
}

func (r *Repository) CheckPhoneNumber(phoneNumber string) error {
	result := r.DB.Where("phone_number = ?", phoneNumber).First(&entity.User{})
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return result.Error
}

func (r *Repository) ConfirmUser(phoneNumber string) error {
	user := entity.User{}
	r.DB.First(&user, "phone_number = ?", phoneNumber)
	if user.CreatedAt.IsZero() {
		return errors.New("user not found")
	}
	user.Confirmed = true
	return r.DB.Save(&user).Error
}

func (r *Repository) GetCredentials(phoneNumber string) (string, error) {
	user := entity.User{}
	result := r.DB.First(&user, "phone_number = ?", phoneNumber)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return "", result.Error
	} else if user.Password == "" || user.Confirmed == false {
		return "", errors.New("user not found")
	}

	return user.Password, nil
}
