package models

import (
	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	gorm.Model
	UserName string `gorm:"size:255;unique" binding:"required"`
	Pwd      string `gorm:"column:pwd;size:255" binding:"required"`
}

// AddUser ...
func AddUser(m *User) error {
	return db.Create(&m).Error
}

// GetUserByUserName ...
func GetUserByUserName(userName string) (v *User, err error) {
	var user User
	if err := db.Where("username = ?", userName).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
