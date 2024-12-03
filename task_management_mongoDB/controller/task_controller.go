package controller

import (
	"context"
	"net/http"
	"task_management_mongoDB/data"
	"task_management_mongoDB/models"
	"time"

	"github.com/gin-gonic/gin"
)

// MongoDB context timeout
const mongoTimeout = 10 * time.Second

// Get all tasks
func GetTasks(ctx *gin.Context) {
	// Set a timeout context for MongoDB operations
	timeoutCtx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()

	tasks, err := data.GetAllTasks(timeoutCtx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// Get a task by ID
func GetTask(ctx *gin.Context) {
	id := ctx.Param("id")

	// Set a timeout context for MongoDB operations
	timeoutCtx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()

	task, err := data.GetTaskByID(timeoutCtx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

// Add a new task
func AddTask(ctx *gin.Context) {
	var newTask models.Task
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set a timeout context for MongoDB operations
	timeoutCtx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()

	err := data.AddTask(timeoutCtx, newTask)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

// Update an existing task
func UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask models.Task
	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set a timeout context for MongoDB operations
	timeoutCtx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()

	err := data.UpdateTask(timeoutCtx, id, updatedTask)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

// Remove a task
func RemoveTask(ctx *gin.Context) {
	id := ctx.Param("id")

	// Set a timeout context for MongoDB operations
	timeoutCtx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()

	err := data.RemoveTask(timeoutCtx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
}
