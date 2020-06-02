package models

import (
	"github.com/jinzhu/gorm"
)

// InitDB 初始化数据库
func InitDB() {
	db := ConnectDB()
	defer db.Close()
	// 自动迁移
	db.AutoMigrate(&File{}, &User{})
}

// ConnectDB 连接数据库
func ConnectDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/arrow?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	// 禁用表名复数
	db.SingularTable(true)
	return db
}
