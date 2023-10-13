package entity

import "time"

type User struct {
	Id          int
	PhoneNumber string `gorm:"unique"`
	Password    string
	Name        string
	Surname     string
	DateOfBirth string
	Url         string    `gorm:"default:null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	Confirmed   bool      `gorm:"default:false"`
}
