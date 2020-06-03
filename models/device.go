package models

import "github.com/jinzhu/gorm"

// Device ...
type Device struct {
	gorm.Model
	UUID      string `gorm:"size(50);unique" binding:"required"`
	Sn        string `gorm:"size(255);unique" binding:"required"`
	Type      string `gorm:"size(255)"`
	Sign      string `gorm:"size(255)"`
	Timestamp int
	Passwd    string `gorm:"size(50)"`
}
