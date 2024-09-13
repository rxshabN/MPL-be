package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
	"github.com/oik17/mpl-be/internal/database"
	"github.com/oik17/mpl-be/internal/models"
)

func GetAllTeamsByScore() ([]models.Teams, error) {
	ctx := context.Background()

	// Check for cached data in Redis
	log.Println("hi")
	cachedData, err := database.RedisClient.Get(ctx, "teams_by_score").Result()
	log.Println(cachedData)
	if cachedData=="" {
		// Cache miss: Fetch data from database
		log.Println("Cache miss: fetching data from database")

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

		// Serialize and cache the data in Redis
		teamsJSON, err := json.Marshal(teams)
		if err != nil {
			return nil, err
		}

		// Set the cache with an expiration of 10 minutes
		err = database.RedisClient.Set(ctx, "teams_by_score", teamsJSON, time.Minute*10).Err()
		if err != nil {
			log.Println("Failed to cache data in Redis:", err)
		} else {
			log.Println("Data cached in Redis successfully")
		}

		return teams, nil
	} else if err != nil {
		log.Println("im here")
		log.Println(err)
		return nil, err
	}

	// Cache hit: Parse the cached JSON data
	log.Println("Cache hit: returning data from Redis")
	var teams []models.Teams
	err = json.Unmarshal([]byte(cachedData), &teams)
	if err != nil {
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

func UpdateTeamHint(teamID string, hint int, remainingTime int) error {
	db := database.DB.Db
	fmt.Printf("%d", remainingTime/1000000000)
	score := (remainingTime / 1000000000) * hint

	fmt.Printf("\n%d", score)
	_, err := db.Exec(`UPDATE team SET hint_number=$1, score =$3 WHERE team_id=$2`, hint, teamID, score)
	if err != nil {
		return err
	}
	return nil
}
