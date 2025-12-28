package routers

import (
	"github.com/gin-gonic/gin"

	"golang-task-management-system/api/controllers"
	"golang-task-management-system/api/middleware"
)

func TaskRoutes(router *gin.Engine) {
	tasks := router.Group("/task")
	{
		tasks.POST("/create", middleware.AuthMiddleware, controllers.CreateTask)
		tasks.POST("/bulkupload", middleware.AuthMiddleware, controllers.CreateTaskBulk)
		tasks.GET("/", controllers.GetTasks)
		tasks.GET("/user/:user_id", middleware.AuthMiddleware, controllers.GetTasksByUserID)
		tasks.PUT("/update/", middleware.AuthMiddleware, controllers.UpdateTask)
		tasks.DELETE("/delete/", middleware.AuthMiddleware, controllers.DeleteTask)
	}
}
