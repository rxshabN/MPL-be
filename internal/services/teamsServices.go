package services

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/oik17/mpl-be/internal/database"
	"github.com/oik17/mpl-be/internal/models"
)

func CreateTeam(team models.Teams) error {
	db := database.DB.Db
	membersJSON, err := json.Marshal(team.TeamMembers)
	if err != nil {
		return err
	}
	_, err = db.Exec(`INSERT INTO teams VALUES $1, $2, $3, $4, $5`, uuid.New(), team.TeamName, membersJSON, team.Score, team.HintNumber)
	if err != nil {
		return err
	}
	return nil
}

func GetAllTeams() ([]models.Teams, error) {
	db := database.DB.Db

	rows, err := db.Query(`SELECT team_id, team_name, team_members, score, hint_number FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []models.Teams

	for rows.Next() {
		var team models.Teams

		err := rows.Scan(&team.TeamID, &team.TeamName, &team.TeamMembers, &team.Score, &team.HintNumber)
		if err != nil {
			return nil, err
		}

		teams = append(teams, team)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return teams, nil
}
