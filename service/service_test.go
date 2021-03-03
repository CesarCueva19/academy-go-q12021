package service

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	cError "github.com/coreos/etcd/error"
	"github.com/go-playground/assert/v2"
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
		t.Run(tt.name, func(t *testing.T) {
			handler := func(w http.ResponseWriter, r *http.Request) {
				io.WriteString(w, string(tt.expectedBody))
			}

			req := httptest.NewRequest("GET", "https://pokeapi.co/api/v2/pokemon/"+tt.id, nil)
			w := httptest.NewRecorder()
			w.WriteHeader(tt.statusCode)
			handler(w, req)

			resp := w.Result()
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(resp.StatusCode)

			assert.Equal(t, tt.statusCode, int(resp.StatusCode))
			assert.Equal(t, tt.expectedBody, body)
		})
	}
}
