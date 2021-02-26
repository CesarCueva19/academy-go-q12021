package router

import (
	"net/http"

	"pokemons/controller"

	"github.com/gorilla/mux"
)

// Controllers - Controllers Interface
type Controllers interface {
	GetPokemons(w http.ResponseWriter, r *http.Request)
	GetPokemonByID(w http.ResponseWriter, r *http.Request)
	CatchPokemons(w http.ResponseWriter, r *http.Request)
}

// New - Creates a router
func New(controllers Controllers) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controller.IndexRoute)
	router.HandleFunc("/pokemons", controllers.GetPokemons).Methods("GET")
	router.HandleFunc("/pokemons/{id}", controllers.GetPokemonByID).Methods("GET")
	router.HandleFunc("/catch/pokemons/{id}", controllers.CatchPokemons).Methods("GET")

	return router
}
