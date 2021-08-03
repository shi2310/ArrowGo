package routers

import (
	"ArrowGo/controllers"
	"ArrowGo/middleware"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "ArrowGo/docs"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors()) // 加入跨域

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/account/register", controllers.Register)
	router.POST("/account/login", controllers.Login)

	router.Use(middleware.JWTAuth()) // 后面的action加入token验证

	router.POST("/user/add", controllers.AddUser)
	router.PUT("/user/changePwd", controllers.ChangePwd)

	return router
}
