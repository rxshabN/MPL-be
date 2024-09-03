package services

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/oik17/mpl-be/internal/database"
	"github.com/oik17/mpl-be/internal/models"
)

func CreateTeam(team models.Teams) error {
	db := database.DB.Db
	_, err := db.Exec(`
        INSERT INTO team (team_id, team_name, team_members, score, hint_number)
        VALUES ($1, $2, $3, $4, $5)`,
		uuid.New(), team.TeamName, pq.Array(team.TeamMembers), team.Score, team.HintNumber)
	if err != nil {
		return err
	}
	return nil
}

func GetAllTeams() ([]models.Teams, error) {
	db := database.DB.Db

	rows, err := db.Query(`SELECT team_id, team_name, team_members, score, hint_number FROM team`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []models.Teams

	for rows.Next() {
		var team models.Teams
		var members pq.StringArray

		err := rows.Scan(&team.TeamID, &team.TeamName, &members, &team.Score, &team.HintNumber)
		if err != nil {
			return nil, err
		}

		team.TeamMembers = []string(members)
		teams = append(teams, team)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return teams, nil
}

func GetAllTeamsByScore() ([]models.Teams, error) {
	db := database.DB.Db

	rows, err := db.Query(`SELECT team_id, team_name, team_members, score, hint_number FROM team ORDER BY score`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []models.Teams

	for rows.Next() {
		var team models.Teams
		var members pq.StringArray

		err := rows.Scan(&team.TeamID, &team.TeamName, &members, &team.Score, &team.HintNumber)
		if err != nil {
			return nil, err
		}

		team.TeamMembers = []string(members)
		teams = append(teams, team)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return teams, nil
}

func UpdateTeamScore(teamID string, score int) error {
	db := database.DB.Db
	_, err := db.Exec(`UPDATE team SET score=$1 WHERE team_id=$2`, score, teamID)
	if err != nil {
		return err
	}
	return nil
}

func GetAllTeamsByHint() ([]models.Teams, error) {
	db := database.DB.Db

	rows, err := db.Query(`SELECT team_id, team_name, team_members, score, hint_number FROM team ORDER BY hint_number`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []models.Teams

	for rows.Next() {
		var team models.Teams
		var members pq.StringArray

		err := rows.Scan(&team.TeamID, &team.TeamName, &members, &team.Score, &team.HintNumber)
		if err != nil {
			return nil, err
		}

		team.TeamMembers = []string(members)
		teams = append(teams, team)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return teams, nil
}

func UpdateTeamHint(teamID string, hint int) error {
	db := database.DB.Db
	_, err := db.Exec(`UPDATE team SET hint_number=$1 WHERE team_id=$2`, hint, teamID)
	if err != nil {
		return err
	}
	return nil
}
