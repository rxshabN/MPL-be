package main

import (
	"net/http"

	"github.com/oik17/mpl-be/internal/database"
	"github.com/oik17/mpl-be/internal/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.Connect()
	database.RedisConnect()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	routes.TeamRoutes(e)
	routes.UserRoutes(e)

	e.Start(":8080")
}
