package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/oik17/mpl-be/internal/controllers"
)

func UserRoutes(e *echo.Echo) {

	e.POST("/signup", controllers.Signup)
	e.POST("/login", controllers.Login)
}
