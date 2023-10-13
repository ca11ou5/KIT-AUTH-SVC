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
	us := entity.User{}
	result := r.DB.Where("phone_number = ?", user.PhoneNumber).First(&us)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		res := r.DB.Create(user)
		return res.Error
	}

	if user.Confirmed == false {
		user.Id = us.Id
		res := r.DB.Save(user)
		return res.Error
	}

	return errors.New("the phone number is already taken")
}
