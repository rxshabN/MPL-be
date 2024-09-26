package services

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/lib/pq"
	"github.com/oik17/mpl-be/internal/database"
	"github.com/oik17/mpl-be/internal/models"
)

func GetAllTeamsByScore() ([]models.Teams, error) {
	ctx := context.Background()

	log.Println("hi")
	cachedData, err := database.RedisClient.Get(ctx, "teams_by_score").Result()
	log.Println(cachedData)
	if cachedData == "" {
		log.Println("Cache miss: fetching data from database")

		db := database.DB.Db
		rows, err := db.Query(`SELECT team_id, team_name, team_members, score, hint_number FROM team ORDER BY score DESC`)
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

		teamsJSON, err := json.Marshal(teams)
		if err != nil {
			return nil, err
		}

		err = database.RedisClient.Set(ctx, "teams_by_score", teamsJSON, time.Minute*5).Err()
		if err != nil {
			log.Println("Failed to cache data in Redis:", err)
		} else {
			log.Println("Data cached in Redis successfully")
		}

		return teams, nil
	} else if err != nil {
		log.Print("Fatal error: ")
		log.Println(err)
		return nil, err
	}

	log.Println("Cache hit: returning data from Redis")
	var teams []models.Teams
	err = json.Unmarshal([]byte(cachedData), &teams)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func UpdateTeamScore(teamID string, scoreToAdd int) error {
	db := database.DB.Db

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var currentScore int
	err = tx.QueryRow(`SELECT score FROM team WHERE team_id=$1`, teamID).Scan(&currentScore)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE team SET score=$1 WHERE team_id=$2`, currentScore+scoreToAdd, teamID)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
func GetAllTeamsByHint() ([]models.Teams, error) {
	db := database.DB.Db

	rows, err := db.Query(`SELECT team_id, team_name, team_members, score, hint_number FROM team ORDER BY hint_number DESC`)
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

func UpdateTeamHint(teamID string, hint int, remainingTime int) (int, error) {
	ctx := context.Background()
	database.RedisClient.Del(ctx, "teams_by_score")
	db := database.DB.Db

	score := (remainingTime / 10000000000) * hint

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var currentScore int
	err = tx.QueryRow(`SELECT score FROM team WHERE team_id=$1`, teamID).Scan(&currentScore)
	if err != nil {
		return 0, err
	}

	_, err = db.Exec(`UPDATE team SET hint_number=$1, score =$3 WHERE team_id=$2`, hint, teamID, score+currentScore)
	if err != nil {
		return score, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return score + currentScore, nil
}
