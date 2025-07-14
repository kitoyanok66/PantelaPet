package main

import (
	"PantelaPet/internal/db"
	"PantelaPet/internal/handlers"
	taskservice "PantelaPet/internal/taskService"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	e := echo.New()

	taskRepo := taskservice.NewTaskRepository(database)
	taskService := taskservice.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", taskHandlers.GetTasks)
	e.POST("/tasks", taskHandlers.PostTasks)
	e.PATCH("/tasks/:id", taskHandlers.PatchTasks)
	e.DELETE("/tasks/:id", taskHandlers.DeleteTasks)

	e.Start("localhost:8080")
}
