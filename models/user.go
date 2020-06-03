package models

import (
	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	gorm.Model
	UserName string `gorm:"size:255;unique" binding:"required"`
	Pwd      string `gorm:"column:pwd;size:255;not null" binding:"required"`
	Photo    string `gorm:"type:text"`
}

// AddUser ...
func AddUser(m *User) error {
	db := ConnectDB()
	defer db.Close()
	return db.Create(&m).Error
}

// GetUserByUserName ...
func GetUserByUserName(userName string) (v *User, err error) {
	var user User
	db := ConnectDB()
	defer db.Close()
	if err := db.Where("user_name = ?", userName).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// ChangePwd ...
func ChangePwd(m *User) error {
	db := ConnectDB()
	defer db.Close()
	return db.Model(m).Update("pwd", m.Pwd).Error
}
