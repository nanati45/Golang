package router

import (
	"task_management_mongoDB/controller"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(client *mongo.Client) *gin.Engine {
	router := gin.Default()

	// Task Routes
	router.GET("/tasks", controller.GetTasks)       // Get all tasks
	router.GET("/tasks/:id", controller.GetTask)    // Get a single task by ID
	router.POST("/tasks", controller.AddTask)       // Add a new task
	router.PUT("/tasks/:id", controller.UpdateTask) // Update an existing task
	router.DELETE("/tasks/:id", controller.RemoveTask) // Remove a task by ID

	return router
}
