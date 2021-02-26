package csv

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"

	"pokemons/model"

	"github.com/sirupsen/logrus"
)

// Service for Pokemon requests
type Service struct {
	logger *logrus.Logger
	csvr   *os.File
	csvw   *csv.Writer
}

// New creates a new PokemonAPI client
func New(
	logger *logrus.Logger,
	csvr *os.File,
	csvw *csv.Writer) *Service {
	return &Service{logger, csvr, csvw}
}

// GetPokemonsInfo - Gets the pokemons info from the csv file
func (s *Service) GetPokemonsInfo() ([]*model.PokemonsData, error) {
	var pokemons []*model.PokemonsData

	csvr := csv.NewReader(s.csvr)
	data, err := csvr.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, row := range data {
		pokemonID, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, err
		}
		pokemonName := row[1]

		pokemonWeight, err := strconv.Atoi(row[2])
		if err != nil {
			return nil, err
		}

		pokemons = append(pokemons, &model.PokemonsData{ID: pokemonID, Name: pokemonName,
			Weight: pokemonWeight})
	}
	s.csvr.Seek(0, 0)

	return pokemons, nil
}

// GetPokemonInfo -  Gets the pokemon info by ID from the csv file
func (s *Service) GetPokemonInfo(pokemonID string) (*model.PokemonsData, error) {
	var pokemons *model.PokemonsData

	csvr := csv.NewReader(s.csvr)
	data, err := csvr.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, row := range data {
		if pokemonID == row[0] {
			rowID, err := strconv.Atoi(row[0])
			if err != nil {
				return nil, err
			}

			pokemonName := row[1]

			pokemonWeight, err := strconv.Atoi(row[2])
			if err != nil {
				return nil, err
			}

			pokemons = &model.PokemonsData{ID: rowID, Name: pokemonName, Weight: pokemonWeight}

			return pokemons, nil
		}
	}
	s.csvr.Seek(0, 0)

	return nil, errors.New("Pokemon not found")
}
