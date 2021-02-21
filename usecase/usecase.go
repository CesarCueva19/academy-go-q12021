package usecase

import (
	"pokemons/models"
	"pokemons/services"
)

// GetPokemons -  Gets the pokemon info by ID from the csv file
func GetPokemons() ([]*models.PokemonsData, error) {
	return services.GetInfo()
}

// GetPokemonByID -  Gets the pokemon info by ID from the csv file
func GetPokemonByID(pokemonID string) (*models.PokemonsData, error) {
	return services.GetPokemonInfo(pokemonID)
}
