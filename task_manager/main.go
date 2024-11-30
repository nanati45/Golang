package main

import (
	"fmt"
	"task_manager/router"
)

func main() {
	fmt.Println("Task Manager API Project")

	r := router.SetupRouter()
	r.Run() // Defaults to ":8080"
}
