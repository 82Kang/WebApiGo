package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/user/RestApi/Models"
	"github.com/user/RestApi/Repository/Players"
	"log"
	"net/http"
	"strconv"
)

func main() {

	router := mux.NewRouter()
	initializeHandlers(router)
	initializeDataStore()
	log.Fatal(http.ListenAndServe(":8000", router))
}
func initializeHandlers(router *mux.Router) {
	router.HandleFunc("/players", GetPlayers).Methods("GET")
	router.HandleFunc("/player/{id}", GetPlayer).Methods("GET")
	router.HandleFunc("/player/{id}", CreatePlayer).Methods("CREATE")
	router.HandleFunc("/player/{id}", DeletePlayer).Methods("DELETE")
}
func GetPlayer(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])
	json.NewEncoder(w).Encode(Players.GetPlayer(i))
}
func GetPlayers(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(Players.Players)
}
func CreatePlayer(w http.ResponseWriter, r *http.Request) {

}
func DeletePlayer(w http.ResponseWriter, r *http.Request) {

}

func initializeDataStore() {
	Players.AddPlayer(Models.Player{Id: 1, Name: "Bayant", ShirtNumber: 11, Address: &Models.Address{Country: "India", State: "Rajasthan", City: "Dhanoor"}})
}
