package main

import (
	"github.com/gin-gonic/gin"

	"golang-task-management-system/api/initializers"
	"golang-task-management-system/api/routers"
)

func init() {
	// LoadEnvVariables()
	initializers.LoadEnvVariables()
	// ConnectToDB()
	initializers.ConnectToDB()
	// SyncDatabase()
	initializers.SyncDatabase()
}

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Golang Task Management System- Hardik Budhwar",
		})
	})

	routers.UserRoutes(r)
	routers.TaskRoutes(r)
	r.Run()
}
