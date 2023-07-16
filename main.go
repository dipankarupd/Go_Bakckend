package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dipankarupd/go-football-betting/pkg/route"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	r := mux.NewRouter()
	route.RegisterGameRoutes(r)
	http.Handle("/", r)

	port := os.Getenv("PORT")

	fmt.Println("Starting the server:")
	log.Fatal(http.ListenAndServe("0.0.0.0"+port, r))
}
