package util

import (
	"math/rand"

	"github.com/dipankarupd/go-football-betting/pkg/model"
)

// CreateRandomGames creates 5 random games with auto-generated game IDs
func CreateRandomGames(num int) []*model.Game {
	var games []*model.Game

	for i := 0; i < num; i++ {
		team1 := generateRandomTeam()
		team2 := generateRandomTeam()

		// Ensure that both teams are unique
		for team1 == team2 {
			team2 = generateRandomTeam()
		}

		game := &model.Game{
			Team1: team1,
			Team2: team2,
		}

		games = append(games, game)
	}

	return games
}

func generateRandomTeam() string {
	teams := []string{"Team A", "Team B", "Team C", "Team D", "Team E"}
	randomIndex := rand.Intn(len(teams))
	return teams[randomIndex]
}
