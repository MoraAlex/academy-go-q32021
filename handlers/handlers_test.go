package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MoraAlex/academy-go-q32021/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var r404 = http.Response{
	Status:     "404 not found",
	StatusCode: 404,
}

var r500 = http.Response{
	Status:     "200",
	StatusCode: 200,
}
var r200 = http.Response{
	Status:     "500 Internal Server Error",
	StatusCode: 500,
}

var p = []*entities.Pokemon{
	{
		Name: "test1",
	},
	{
		Name: "test2",
	},
	{
		Name: "test3",
	},
}

type mockGetPokemonUseCase struct {
	mock.Mock
}

type mockGetPokemonsUseCase struct {
	mock.Mock
}

func (m mockGetPokemonUseCase) GetPokemon(id string) (*entities.Pokemon, error) {
	args := m.Called(id)
	return args[0].(*entities.Pokemon), args.Error(1)
}

func (m mockGetPokemonsUseCase) GetPokemons() ([]*entities.Pokemon, error) {
	args := m.Called()
	return args[0].([]*entities.Pokemon), args.Error(1)
}

func TestGetAllPokemons(t *testing.T) {
	testCases := []struct {
		name         string
		response     []*entities.Pokemon
		httpResponse http.Response
		hasError     bool
		err          error
		url          string
	}{
		{
			name:         "No error",
			response:     p,
			httpResponse: r200,
			hasError:     false,
			err:          nil,
			url:          "/pokemons",
		},
		{
			name:         "Error",
			response:     nil,
			httpResponse: r500,
			hasError:     true,
			err:          errors.New("test"),
			url:          "/pokemons",
		},
		{
			name:         "Not found",
			response:     nil,
			httpResponse: r404,
			hasError:     true,
			err:          errors.New("test"),
			url:          "/pokmons",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mgp := mockGetPokemonsUseCase{}
			mgp.On("GetPokemons").Return(tc.response, tc.err)
			mgps := mockGetPokemonUseCase{}
			h := New(mgp, mgps)
			r := httptest.NewRequest(http.MethodGet, tc.url, nil)
			w := httptest.NewRecorder()
			h.GetAllPokemons(w, r)
			res := w.Result()
			//data, err := ioutil.ReadAll(res.Body)
			//.Println(res, data, err)
			if tc.hasError {
				assert.EqualValues(t, res.StatusCode, tc.httpResponse.StatusCode)
			}
		})
	}
}

func TestGetPokemon(t *testing.T) {
	testCases := []struct {
		name         string
		response     []*entities.Pokemon
		httpResponse http.Response
		hasError     bool
		err          error
		url          string
		id           string
	}{
		{
			name:         "No error",
			response:     p,
			httpResponse: r200,
			hasError:     false,
			err:          nil,
			url:          "/pokemons/1",
			id:           "1",
		},
		{
			name:         "Error",
			response:     nil,
			httpResponse: r500,
			hasError:     true,
			err:          errors.New("test"),
			url:          "/pokemons/",
			id:           "",
		},
		{
			name:         "Not found",
			response:     nil,
			httpResponse: r404,
			hasError:     true,
			err:          errors.New("test"),
			url:          "/pokmons/a",
			id:           "a",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mgp := mockGetPokemonsUseCase{}
			mgps := mockGetPokemonUseCase{}
			mgps.On("GetPokemon", tc.id).Return(tc.response, tc.err)
			h := New(mgp, mgps)
			r := httptest.NewRequest(http.MethodGet, tc.url, nil)
			w := httptest.NewRecorder()
			h.GetPokemon(w, r)
			res := w.Result()
			//data, err := ioutil.ReadAll(res.Body)
			//fmt.Println(res, data, err)
			if tc.hasError {
				assert.EqualValues(t, res.StatusCode, tc.httpResponse.StatusCode)
			}
		})
	}
}
