package service

import (
	"net/http"
	"testing"
	"time"

	cError "github.com/coreos/etcd/error"
	"github.com/go-playground/assert/v2"
	"gopkg.in/resty.v1"
)

func Test_PokemonAPI(t *testing.T) {
	test := []struct {
		name         string
		statusCode   int
		expectedBody []byte
		id           string
		err          error
	}{
		{
			name:       "Pokemon found, response OK",
			statusCode: http.StatusOK,
			expectedBody: []byte(
				`{
				"id": 1,
				"name": "bulbasaur",
				}`,
			),
			id:  "1",
			err: nil,
		},
		{
			name:       "Pokemon not found, response Not found",
			statusCode: http.StatusNotFound,
			expectedBody: []byte(
				`Not Found`,
			),
			id:  "100000",
			err: cError.NewError(http.StatusNotFound, "error", 0),
		},
	}

	for _, tt := range test {
		client := resty.New().
			SetHostURL("https://pokeapi.co/api/v2/").
			SetTimeout(time.Second * 10).
			OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
				if r.IsSuccess() {
					return nil
				}

				return cError.NewError(r.StatusCode(), "error", 0)
			})

		s := &Service{
			client: client,
		}

		t.Run(tt.name, func(t *testing.T) {
			resp, err := s.client.R().
				SetPathParams(map[string]string{"pokemonId": tt.id}).
				SetHeader("Accept", "application/json").
				Get("pokemon/{pokemonId}")

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.statusCode, int(resp.StatusCode()))

		})
	}
}
