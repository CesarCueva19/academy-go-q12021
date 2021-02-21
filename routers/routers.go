package routers

import (
	"pokemons/controllers"

	"github.com/gorilla/mux"
)

// NewRouters - Creates a router
func NewRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controllers.IndexRoute)
	router.HandleFunc("/pokemons", controllers.GetPokemons).Methods("GET")
	router.HandleFunc("/pokemons/{id}", controllers.GetPokemonByID).Methods("GET")

	return router
}
