package controller

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

// Get all tasks
func GetTasks(ctx *gin.Context) {
	tasks := data.GetAllTasks()
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// Get a task by ID
func GetTask(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := data.GetTaskByID(id)
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
	data.AddTask(newTask)
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
	err := data.UpdateTask(id, updatedTask)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

// Remove a task
func RemoveTask(ctx *gin.Context) {
	id := ctx.Param("id")
	err := data.RemoveTask(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
}
