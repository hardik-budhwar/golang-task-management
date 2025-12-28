package routers

import (
	"github.com/gin-gonic/gin"

	"golang-task-management-system/api/controllers"
	"golang-task-management-system/api/middleware"
)

func UserRoutes(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.POST("/register", controllers.UserRegistration)
		user.POST("/login", controllers.UserLogin)
		user.PUT("/logout", middleware.AuthMiddleware, controllers.UserLogout)
		user.DELETE("/delete", middleware.AuthMiddleware, controllers.UserDeletion)
		user.GET("/", middleware.AuthMiddleware, controllers.GetAllUsers)
	}
}
