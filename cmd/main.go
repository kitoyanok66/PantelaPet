package main

import (
	"PantelaPet/internal/db"
	"PantelaPet/internal/handlers"
	taskservice "PantelaPet/internal/taskService"
	userservice "PantelaPet/internal/userService"
	"PantelaPet/internal/web/tasks"
	"PantelaPet/internal/web/users"
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

	userRepo := userservice.NewUserRepository(database)
	userService := userservice.NewUserService(userRepo)
	userHandlers := handlers.NewUserHandler(userService)

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	strictTaskHandler := tasks.NewStrictHandler(taskHandlers, nil)
	tasks.RegisterHandlers(e, strictTaskHandler)

	strictUserHandler := users.NewStrictHandler(userHandlers, nil)
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
