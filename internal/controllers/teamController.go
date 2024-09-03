package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/oik17/mpl-be/internal/models"
	"github.com/oik17/mpl-be/internal/services"
)

func CreateTeam(c echo.Context) error {
	var team models.Teams
	if err := c.Bind(&team); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to create team",
			"data":    err.Error(),
		})
	}

	err := services.CreateTeam(team)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to create team",
			"data":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully created team",
		"data":    team.TeamName,
	})
}

func GetAllTeams(c echo.Context) error {
	teams, err := services.GetAllTeams()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to fetch teams",
			"data":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, teams)
}

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
	err := services.UpdateTeamHint(teamID, hint.Hint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to update team score",
			"data":    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully updated team score",
		"data":    strconv.Itoa(hint.Hint),
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
