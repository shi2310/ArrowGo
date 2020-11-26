package models

import "github.com/jinzhu/gorm"

// UserDevice 用户设备绑定信息
type UserDevice struct {
	gorm.Model
	Mid int    //User.ID
	Sn  string `gorm:"size(255)"` // Device.Sn
}
