package models

import (
	"github.com/jinzhu/gorm"
)

// File 文件信息模型
type File struct {
	gorm.Model
	Name    string `gorm:"size(255)"`
	URL     string `gorm:"size(255)"`
	FileMd5 string `gorm:"column:filemd5;size(255)"`
}

// AddFile ...
func AddFile(m *File) error {
	db := ConnectDB()
	defer db.Close()
	return db.Create(&m).Error
}

// GetFileByMD5 ...
func GetFileByMD5(md5 string) (v *File, err error) {
	var file File
	db := ConnectDB()
	defer db.Close()
	if err := db.Where("filemd5 = ?", md5).First(&file).Error; err != nil {
		return nil, err
	}
	return &file, nil
}
