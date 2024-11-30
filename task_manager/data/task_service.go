package data

import (
	"errors"
	"task_manager/models"
)

var tasks = []models.Task{
	{ID: "1", Title: "Task Manager Project", Description: "Add/View/Delete Tasks", DueDate: models.CurrentTime(), Status: "In Progress"},
	{ID: "2", Title: "Books Management Project", Description: "Add/View/Delete Books", DueDate: models.CurrentTime().AddDate(0, 0, -1), Status: "Completed"},
}

// Get all tasks
func GetAllTasks() []models.Task {
	return tasks
}

// Get a task by ID
func GetTaskByID(id string) (models.Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

// Add a task
func AddTask(task models.Task) {
	tasks = append(tasks, task)
}

// Update a task
func UpdateTask(id string, updatedTask models.Task) error {
	for i, task := range tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			if !updatedTask.DueDate.IsZero() {
				tasks[i].DueDate = updatedTask.DueDate
			}
			if updatedTask.Status != "" {
				tasks[i].Status = updatedTask.Status
			}
			return nil
		}
	}
	return errors.New("task not found")
}

// Remove a task
func RemoveTask(id string) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
