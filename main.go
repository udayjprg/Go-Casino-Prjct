package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type player struct {
	Id        int    `json:"Id"`
	Name      string `json:"Name"`
	AccountNo int    `json:"AccountNo"`
	Balance   int    `json:"Balance"`
}

func handlePlayer() {
	r := mux.NewRouter()
	r.HandleFunc("/createplayer", createPlayer())
	r.HandleFunc("/getplayers", getPlayers)
	r.HandleFunc("/getbalance", getBalance)
	log.Fatal(http.ListenAndServe(":80080", r))
}

var players = []player{
	{001, "Uday", 22335, 1200},
	{002, "Jon", 3456, 2500},
	{003, "Steve", 55689, 5000},
}

func getPlayers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered getPlayer function")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}
func createPlayer(w http.ResponseWriter, r *http.Request) {

}
