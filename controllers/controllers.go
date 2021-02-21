package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pokemons/usecase"

	"github.com/gorilla/mux"
)

// IndexRoute - This is the controller for default route
func IndexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Pokemons API")
}

// GetPokemons - This is the controller for getting the pokemons info
func GetPokemons(w http.ResponseWriter, r *http.Request) {
	pokemons, err := usecase.GetPokemons()
	if err != nil {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemons)
}

// GetPokemonByID - This is the controller for getting the pokemon info by ID
func GetPokemonByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pokemon, err := usecase.GetPokemonByID(params["id"])
	if err != nil {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemon)
}
