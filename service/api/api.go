package api

import (
	"encoding/csv"
	"encoding/json"
	"strconv"
	"time"

	"pokemons/model"

	cError "github.com/coreos/etcd/error"
	"github.com/sirupsen/logrus"
	"gopkg.in/resty.v1"
)

// Service for Pokemon requests
type Service struct {
	logger *logrus.Logger
	client *resty.Client
	csvw   *csv.Writer
}

const pokemonEndpoint = "pokemon/{pokemon_id}"

// New creates a new PokemonAPI client
func New(
	logger *logrus.Logger,
	host string,
	timeout time.Duration,
	csvw *csv.Writer) (*Service, error) {
	client := resty.New().
		SetHostURL(host).
		SetTimeout(timeout).
		OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
			if r.IsSuccess() {
				return nil
			}

			return cError.NewError(r.StatusCode(), "error", 0)
		})

	return &Service{logger, client, csvw}, nil
}

// PokemonAPI - Handles the client
func (s *Service) PokemonAPI(pokemonID string) (*model.PokemonsData, error) {
	var pokeAPI *model.PokemonsData

	resp, err := s.client.R().
		SetPathParams(map[string]string{"pokemon_id": pokemonID}).
		SetHeader("Accept", "application/json").
		Get(pokemonEndpoint)
	if err != nil {
		return nil, err
	}

	// Unmarshal Pokemon model
	if err := json.Unmarshal(resp.Body(), &pokeAPI); err != nil {
		return nil, err
	}

	if err := s.csvw.Write([]string{strconv.Itoa(pokeAPI.ID), pokeAPI.Name,
		strconv.Itoa(pokeAPI.Weight)}); err != nil {
		return nil, err
	}

	s.csvw.Flush()

	return pokeAPI, nil
}
