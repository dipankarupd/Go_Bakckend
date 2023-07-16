package util

import (
	"math/rand"
	"time"
)

func PredictWinner(team1 string, team2 string, id int) string {

	rand.Seed(time.Now().UnixNano())

	// Generate a random integer between 0 and 100
	randomInt := rand.Intn(10001)

	if (randomInt+id)%3 == 0 {
		return team1
	} else {
		return team2
	}
}
