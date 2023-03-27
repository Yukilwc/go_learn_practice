package initialize

import (
	"gjm/middleware"

	userController "gjm/controller/user"

	"github.com/gin-gonic/gin"
)

// 初始化gin
func InitRouter() {
	// Router := gin.New()
	// Router.Use(ginzap.Ginzap(global.LOG, time.RFC3339, true))
	// Router.Use(ginzap.RecoveryWithZap(global.LOG, true))
	Router := gin.Default()
	// 跨域
	Router.Use(middleware.Cors())
	// authGroup := Router.Group("/")
	// {
	// 	authGroup.POST("/list", ListController)
	// }
	// authGroup.Use(JWTAuth())
	Router.POST("/list", middleware.JWTAuth(), userController.ListController)

	publicGroup := Router.Group("/")
	{
		publicGroup.POST("register", userController.RegisterController)
		publicGroup.POST("login", userController.LoginController)
	}
	Router.Run()
}
