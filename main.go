package pantelapet

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type requestBody struct {
	Task string `json:"Task"`
}

var task string

func getTask(c echo.Context) error {
	return c.JSON(http.StatusOK, task)
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/task", getTask)

	e.Start("localhost:8080")
}
