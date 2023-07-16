package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dipankarupd/go-football-betting/pkg/route"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	r := mux.NewRouter()
	route.RegisterGameRoutes(r)
	http.Handle("/", r)

	fmt.Println("Starting the server at port 8080:")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
