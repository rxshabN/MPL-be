package main

import (
	"fmt"
	"github.com/oik17/mpl-be/internal/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	database.Connect()
	e := echo.New()
	fmt.Println("hello")
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, Echo!"})
	})
	e.Start(":8080")
}
