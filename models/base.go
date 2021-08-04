package models

import (
	"github.com/jinzhu/gorm"
)

// 在其它model的实体类中可直接调用
var db *gorm.DB

// InitDB 初始化数据库
func InitDB() *gorm.DB {
	_db, err := gorm.Open("mysql", "root:mr123321@tcp(192.168.1.10:3306)/arrow?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		_db.Close()
	}
	// 设置连接池，空闲连接
	_db.DB().SetMaxIdleConns(50)
	// 打开链接
	_db.DB().SetMaxOpenConns(100)
	// 禁用表名复数
	_db.SingularTable(true)
	// 自动迁移
	_db.AutoMigrate(&File{}, &User{})
	db = _db
	return db
}
