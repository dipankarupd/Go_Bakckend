package controller

import (
	"encoding/json"
	"net/http"

	"github.com/dipankarupd/go-football-betting/pkg/model"
	"github.com/dipankarupd/go-football-betting/pkg/util"
)

type GameResponse struct {
	ID    uint   `json:"id"`
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
}

type PredictResponse struct {
	ID     uint   `json:"id"`
	Winner string `json:"winner"`
}

type GameRequest struct {
	ID uint `json:"id"`
}

var NewGame model.Game

func GetGames(w http.ResponseWriter, r *http.Request) {
	games := util.CreateRandomGames(5)

	for _, value := range games {
		model.CreateGame(value)
	}

	newGames := model.GetGames()
	gameResponses := make([]GameResponse, len(newGames))

	for i, game := range newGames {
		gameResponses[i] = GameResponse{
			ID:    game.ID,
			Team1: game.Team1,
			Team2: game.Team2,
		}
	}

	res, err := json.Marshal(gameResponses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func PredictWinner(w http.ResponseWriter, r *http.Request) {
	gameRequest := &GameRequest{}
	err := json.NewDecoder(r.Body).Decode(gameRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := gameRequest.ID
	gameDetail, _ := model.GetGameById(int64(id))

	predictedValue := PredictResponse{
		ID:     gameDetail.ID,
		Winner: util.PredictWinner(gameDetail.Team1, gameDetail.Team2, int(gameDetail.ID)),
	}
	model.UpdateDb(int(predictedValue.ID), predictedValue.Winner)
	res, err := json.Marshal(predictedValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
