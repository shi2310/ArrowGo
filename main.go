package main

import (
	"ArrowGo/models"
	"ArrowGo/routers"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// 初始化路由(使用Gin)
	r := routers.InitRouter()

	// 初始化数据models(使用Gorm)
	db := models.InitDB()
	defer db.Close()

	if err := r.Run(":8044"); err != nil {
		log.Fatal("程序启动失败:", err)
	}
}
