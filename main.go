package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Task struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type TaskRequest struct {
	Name string `json:"name"`
}

var tasks = []Task{}

func getTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

func postTasks(c echo.Context) error {
	var req TaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	task := Task{
		ID:     uuid.New().String(),
		Name:   req.Name,
		Status: "In process",
	}
	tasks = append(tasks, task)
	return c.JSON(http.StatusCreated, task)
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", getTasks)
	e.POST("/tasks", postTasks)

	e.Start("localhost:8080")
}
