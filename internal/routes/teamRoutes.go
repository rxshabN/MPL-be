package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/oik17/mpl-be/internal/controllers"
)

func TeamRoutes(e *echo.Echo) {
	r := e.Group("/teams")

	r.POST("/createTeam", controllers.CreateTeam)
	r.GET("/getAllTeams", controllers.GetAllTeams)
}
