package repository

import (
	"Auth_Service/configs"
	"Auth_Service/internal/entity"
	"errors"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func InitRepository(cfg *configs.Config) Repository {
	return Repository{DB: InitDB(cfg)}
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
