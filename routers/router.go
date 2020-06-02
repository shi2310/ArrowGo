package routers

import (
	"ArrowGo/controllers"
	"ArrowGo/middleware"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors()) // 加入跨域

	router.POST("/user/register", controllers.Register)
	router.POST("/user/login", controllers.Login)
	router.Use(middleware.JWTAuth()) // 后面的action加入token验证
	router.POST("/user/add", controllers.AddUser)
	router.POST("/user/changePwd", controllers.ChangePwd)
	return router
}
