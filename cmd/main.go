package main

import (
	"net/http"

	"github.com/oik17/mpl-be/internal/database"
	"github.com/oik17/mpl-be/internal/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	database.Connect()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, Echo!"})
	})
	routes.TeamRoutes(e)
	e.Start(":8080")
}
