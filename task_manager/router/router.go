package router

import (
	"task_manager/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Task Routes
	router.GET("/tasks", controller.GetTasks)
	router.GET("/tasks/:id", controller.GetTask)
	router.DELETE("/tasks/:id", controller.RemoveTask)
	router.PUT("/tasks/:id", controller.UpdateTask)
	router.POST("/tasks", controller.AddTask)

	return router
}
