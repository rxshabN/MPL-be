package services

import (
	"github.com/google/uuid"
	"github.com/oik17/mpl-be/internal/database"
	"github.com/oik17/mpl-be/internal/models"
)

func CreateTeam(team models.Teams) error {
	db := database.DB.Db
	_, err := db.Exec(`INSERT INTO teams VALUES $1, $2, $3, $4, $5`, uuid.New(), team.TeamName, team.TeamMembers, team.Score, team.HintNumber)
	if err != nil {
		return err
	}
	return nil
}
