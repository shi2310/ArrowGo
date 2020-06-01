package main

import (
	"ArrowGo/models"
	"ArrowGo/routers"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	// 初始化数据models(使用Gorm)
	models.InitDB()
}

func main() {
	// 初始化路由(使用Gin)
	r := routers.InitRouter()
	r.Run(":8044")
}
