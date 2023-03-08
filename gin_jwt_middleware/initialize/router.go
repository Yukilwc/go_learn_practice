package initialize

import (
	"gjm/middleware"

	userController "gjm/controller/user"

	"github.com/gin-gonic/gin"
)

// 初始化gin
func InitRouter() {
	Router := gin.Default()
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
