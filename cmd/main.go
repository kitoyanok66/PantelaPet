package main

import (
	"PantelaPet/internal/db"
	"PantelaPet/internal/handlers"
	taskservice "PantelaPet/internal/taskService"
	"PantelaPet/internal/web/tasks"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	taskRepo := taskservice.NewTaskRepository(database)
	taskService := taskservice.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskService)

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	strictHandler := tasks.NewStrictHandler(taskHandlers, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
