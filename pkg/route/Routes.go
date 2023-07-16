package route

import (
	"github.com/dipankarupd/go-football-betting/pkg/controller"
	"github.com/gorilla/mux"
)

var RegisterGameRoutes = func(router *mux.Router) {
	router.HandleFunc("/games/", controller.PredictWinner).Methods("POST")
	router.HandleFunc("/games/", controller.GetGames).Methods("GET")
}
