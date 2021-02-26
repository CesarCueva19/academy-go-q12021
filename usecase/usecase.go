package usecase

import (
	"pokemons/model"
)

// Usecase - Usecase model
type Usecase struct {
	pokemonService PokemonService
	csvService     CSVService
}

// PokemonService - Service interface
type PokemonService interface {
	PokemonAPI(pokemonID string) (*model.PokemonsData, error)
}

// CSVService - Service interface
type CSVService interface {
	GetPokemonsInfo() ([]*model.PokemonsData, error)
	GetPokemonInfo(pokemonID string) (*model.PokemonsData, error)
}

// New - Create Usecase
func New(pokemonService PokemonService, csvService CSVService) *Usecase {
	return &Usecase{pokemonService, csvService}
}

// GetPokemons -  Gets the pokemons info from the csv file
func (u *Usecase) GetPokemons() ([]*model.PokemonsData, error) {
	return u.csvService.GetPokemonsInfo()
}

// GetPokemonByID -  Gets the pokemon info by ID from the csv file
func (u *Usecase) GetPokemonByID(pokemonID string) (*model.PokemonsData, error) {
	return u.csvService.GetPokemonInfo(pokemonID)
}

// InsertPokemonByID -  Inserts the pokemon info by ID to the csv file from PokemonAPI
func (u *Usecase) InsertPokemonByID(pokemonID string) (*model.PokemonsData, error) {
	return u.pokemonService.PokemonAPI(pokemonID)
}
