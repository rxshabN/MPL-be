package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/oik17/mpl-be/internal/services"
	"github.com/oik17/mpl-be/internal/utils"
)

func UpdateTeamScore(c echo.Context) error {
	teamID := c.Param("teamID")
	var score struct {
		Score int `json:"score"`
	}
	if err := c.Bind(&score); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid input",
			"data":    err.Error(),
		})
	}

	err := services.UpdateTeamScore(teamID, score.Score)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to update team score",
			"data":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully updated team score",
		"data":    strconv.Itoa(score.Score),
	})
}

func GetAllTeamsByScore(c echo.Context) error {
	teams, err := services.GetAllTeamsByScore()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to fetch teams",
			"data":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, teams)
}

func UpdateTeamHint(c echo.Context) error {
	
	teamID := c.Param("teamID")
	var hint struct {
		Hint int `json:"hint"`
	}

	if err := c.Bind(&hint); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid input",
			"data":    err.Error(),
		})
	}
	if utils.GlobalTimer == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Timer has not been started yet",
		})
	}

	remainingTime := utils.GlobalTimer.TimeLeft()

	score, err := services.UpdateTeamHint(teamID, hint.Hint, remainingTime)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to update team score",
			"data":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully updated team score",
		"data":    strconv.Itoa(score),
	})
}

func GetAllTeamsByHints(c echo.Context) error {
	teams, err := services.GetAllTeamsByHint()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to fetch teams",
			"data":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, teams)
}
func StartTimer(c echo.Context) error {
	utils.CreateTimer(time.Hour)

	log.Println("Timer started")
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully started timer",
	})
}

func GetTimeLeft(c echo.Context) error {
	if utils.GlobalTimer == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Timer has not been started yet",
		})
	}

	remainingTime := utils.GlobalTimer.TimeLeft()

	return c.JSON(http.StatusOK, map[string]int{
		"time_left": remainingTime / 1000000000,
	})
}
