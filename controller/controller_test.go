package controller

import (
	"errors"
	"testing"

	"pokemons/controller/mocks"
	"pokemons/model"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/unrolled/render"
)

//mockgen -source=controller/controller.go -destination=controller/mocks/controller.go -package=mocks

var (
	pokemon1 = model.PokemonsData{
		ID:     1,
		Name:   "bulbasaur",
		Weight: 69,
	}
	pokemon2 = model.PokemonsData{
		ID:     2,
		Name:   "ivysaur",
		Weight: 130,
	}
	successPokemons = []*model.PokemonsData{&pokemon1, &pokemon2}
)

func Test_GetPokemons(t *testing.T) {
	l := &logrus.Logger{}
	r := render.New()
	tests := []struct {
		name                    string
		expectedUsecaseResponse []*model.PokemonsData
		expectedError           error
		wantError               bool
	}{
		{
			name:                    "OK, get pokemon",
			expectedUsecaseResponse: successPokemons,
			wantError:               false,
			expectedError:           nil,
		},
		{
			name:          "Pokemon Not Found",
			wantError:     true,
			expectedError: errors.New("Not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			u := mocks.NewMockUsecase(mockCtrl)

			u.EXPECT().GetPokemons().Return(tt.expectedUsecaseResponse, tt.expectedError)

			c := New(u, l, r)
			response, err := c.usecase.GetPokemons()

			assert.Equal(t, response, tt.expectedUsecaseResponse)

			if tt.wantError {
				assert.NotNil(t, err)
				assert.Equal(t, err, tt.expectedError)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func Test_GetPokemonByID(t *testing.T) {
	l := &logrus.Logger{}
	r := render.New()
	tests := []struct {
		name                    string
		expectedParams          string
		expectedUsecaseResponse *model.PokemonsData
		expectedError           error
		wantError               bool
	}{
		{
			name:                    "OK, get pokemon",
			expectedParams:          "1",
			expectedUsecaseResponse: &pokemon1,
			wantError:               false,
			expectedError:           nil,
		},
		{
			name:           "Pokemon Not Catched",
			expectedParams: "1000",
			wantError:      true,
			expectedError:  errors.New("The pokemon has not been catched"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			u := mocks.NewMockUsecase(mockCtrl)

			u.EXPECT().GetPokemonByID(tt.expectedParams).Return(tt.expectedUsecaseResponse, tt.expectedError)

			c := New(u, l, r)
			response, err := c.usecase.GetPokemonByID(tt.expectedParams)

			assert.Equal(t, response, tt.expectedUsecaseResponse)

			if tt.wantError {
				assert.NotNil(t, err)
				assert.Equal(t, err, tt.expectedError)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func Test_CatchPokemons(t *testing.T) {
	l := &logrus.Logger{}
	r := render.New()
	tests := []struct {
		name                    string
		expectedParams          string
		expectedUsecaseResponse *model.PokemonsData
		expectedError           error
		wantError               bool
	}{
		{
			name:                    "OK,pokemon catched",
			expectedParams:          "1",
			expectedUsecaseResponse: &pokemon1,
			wantError:               false,
			expectedError:           nil,
		},
		{
			name:           "Pokemon Not Catched",
			expectedParams: "1000",
			wantError:      true,
			expectedError:  errors.New("The pokemon could not be catched"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			u := mocks.NewMockUsecase(mockCtrl)

			u.EXPECT().InsertPokemonByID(tt.expectedParams).Return(tt.expectedUsecaseResponse, tt.expectedError)

			c := New(u, l, r)
			response, err := c.usecase.InsertPokemonByID(tt.expectedParams)

			assert.Equal(t, response, tt.expectedUsecaseResponse)

			if tt.wantError {
				assert.NotNil(t, err)
				assert.Equal(t, err, tt.expectedError)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
