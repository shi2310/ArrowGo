package models

import "github.com/jinzhu/gorm"

var db *gorm.DB
var err error

// InitDB 初始化数据库
func InitDB() {
	db, err = gorm.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/arrow?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SingularTable(true)
	// 自动迁移
	db.AutoMigrate(&File{}, &User{})
}
