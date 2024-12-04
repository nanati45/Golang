package router

import (
	"task_management_mongoDB/controller"
	"task_management_mongoDB/middleware" // Import your middleware package

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(client *mongo.Client) *gin.Engine {
	router := gin.Default()

	// Public Routes (no authentication required)
	router.POST("/register", controller.RegisterUser) // User registration
	router.POST("/login", controller.LoginUser)       // User login

	// Protected Routes (authentication required)
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware()) // Apply JWT authentication middleware to all routes in this group

	// Task Routes
	protected.GET("/tasks", controller.GetTasks)       // Get all tasks
	protected.GET("/tasks/:id", controller.GetTask)    // Get a single task by ID
	protected.POST("/tasks", controller.AddTask)       // Add a new task
	protected.PUT("/tasks/:id", controller.UpdateTask) // Update an existing task
	protected.DELETE("/tasks/:id", controller.RemoveTask) // Remove a task by ID

	return router
}
