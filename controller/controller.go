package controller

import (
	"fmt"
	"net/http"

	"pokemons/model"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

// PokemonController - PokemonController model
type PokemonController struct {
	usecase Usecase
	logger  *logrus.Logger
	render  *render.Render
}

// Usecase - Usecase interface
type Usecase interface {
	GetPokemons() ([]*model.PokemonsData, error)
	GetPokemonByID(pokemonID string) (*model.PokemonsData, error)
	InsertPokemonByID(pokemonID string) (*model.PokemonsData, error)
}

// New - Creates Controller
func New(
	u Usecase,
	log *logrus.Logger,
	r *render.Render,
) *PokemonController {
	return &PokemonController{u, log, r}
}

// IndexRoute - This is the controller for default route
func IndexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Pokemons API")
}

// GetPokemons - This is the controller for getting the pokemons info
func (p *PokemonController) GetPokemons(w http.ResponseWriter, r *http.Request) {
	pokemons, err := p.usecase.GetPokemons()
	if err != nil {
		p.logger.Errorf("Getting Pokemons %v", err)
		p.render.JSON(w, http.StatusNotFound, "Errot Getting Pokemons")

		return
	}
	p.render.JSON(w, http.StatusOK, pokemons)
}

// GetPokemonByID - This is the controller for getting the pokemon info by ID
func (p *PokemonController) GetPokemonByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pokemon, err := p.usecase.GetPokemonByID(params["id"])
	if err != nil {
		p.logger.Errorf("Getting Pokemon %v", err)
		p.render.JSON(w, http.StatusNotFound, "Error Getting Pokemon")
		return
	}
	p.render.JSON(w, http.StatusOK, pokemon)
}

// CatchPokemons - This is the controller for getting the pokemon from the API
func (p *PokemonController) CatchPokemons(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pokemon, err := p.usecase.InsertPokemonByID(params["id"])
	if err != nil {
		p.logger.Errorf("Getting Pokemon %v", err)
		p.render.JSON(w, http.StatusNotFound, "Error Getting Pokemon")
		return
	}
	p.render.JSON(w, http.StatusOK, pokemon)
}
