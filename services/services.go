package services

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"

	"pokemons/models"
)

// GetInfo - Gets the pokemons info from the csv file
func GetInfo() ([]*models.PokemonsData, error) {
	var pokemons []*models.PokemonsData

	//TODO: Move this logic to a config file
	file, err := os.Open("pokedex.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	read := csv.NewReader(file)

	for {
		storaged, err := read.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			if errorParsed, ok := err.(*csv.ParseError); ok {
				if errorParsed == csv.ErrFieldCount {
					continue
				}
			}
		}

		pokemonID, err := strconv.Atoi(storaged[0])
		if err != nil {
			return nil, err
		}
		pokemonName := storaged[1]

		pokemons = append(pokemons, &models.PokemonsData{ID: pokemonID, Name: pokemonName, Type: storaged[2]})
	}
	return pokemons, nil
}

// GetPokemonInfo -  Gets the pokemon info by ID from the csv file
func GetPokemonInfo(pokemonID string) (*models.PokemonsData, error) {

	file, err := os.Open("pokedex.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	read := csv.NewReader(file)

	for {
		storaged, err := read.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			if errorParsed, ok := err.(*csv.ParseError); ok {
				if errorParsed == csv.ErrFieldCount {
					continue
				}
			}
		}

		if pokemonID == storaged[0] {
			rowID, err := strconv.Atoi(storaged[0])
			if err != nil {
				return nil, err
			}
			return &models.PokemonsData{ID: rowID, Name: storaged[1], Type: storaged[2]}, nil
		}
	}

	return nil, errors.New("Pokemon not found")
}
