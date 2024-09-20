package routes

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/oik17/mpl-be/internal/controllers"
	"github.com/oik17/mpl-be/internal/middleware"
)

func TeamRoutes(e *echo.Echo) {
	r := e.Group("/teams")

	r.POST("/createTeam", controllers.CreateTeam, echojwt.JWT(controllers.JWTSecret), middleware.Protected)
	r.GET("/getAllTeams", controllers.GetAllTeams)

	r.GET("/getTeamsByScore", controllers.GetAllTeamsByScore)
	r.PUT("/updateScore/:teamID", controllers.UpdateTeamScore, echojwt.JWT(controllers.JWTSecret), middleware.Protected)

	r.GET("/getTeamsByHint", controllers.GetAllTeamsByHints)
	r.PUT("/updateHint/:teamID", controllers.UpdateTeamHint, echojwt.JWT(controllers.JWTSecret), middleware.Protected)

	r.GET("/getTeamHint/:teamID", controllers.GetTeamHint, echojwt.JWT(controllers.JWTSecret), middleware.Protected)

	r.GET("/startTimer", controllers.StartTimer, echojwt.JWT(controllers.JWTSecret), middleware.Protected)
	r.GET("/getTimer", controllers.GetTimeLeft)
}
