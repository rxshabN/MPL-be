package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/oik17/mpl-be/internal/controllers"
)

func TeamRoutes(e *echo.Echo) {
	r := e.Group("/teams")

	r.POST("/createTeam", controllers.CreateTeam)
	r.GET("/getAllTeams", controllers.GetAllTeams)

	r.GET("/getTeamsByScore", controllers.GetAllTeamsByScore)
	r.PUT("/updateScore/:teamID", controllers.UpdateTeamScore)

	r.GET("/getTeamsByHint", controllers.GetAllTeamsByHints)
	r.PUT("/updateHint/:teamID", controllers.UpdateTeamHint)

	r.GET("/getTeamHint/:teamID", controllers.GetTeamHint)

	r.GET("/startTimer", controllers.StartTimer)
	r.GET("/getTimer", controllers.GetTimeLeft)
}
