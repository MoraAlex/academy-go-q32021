package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MoraAlex/academy-go-q32021/entities"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var r400 = http.Response{
	Status:     http.StatusText(http.StatusBadRequest),
	StatusCode: http.StatusBadRequest,
}

var r200 = http.Response{
	Status:     http.StatusText(http.StatusOK),
	StatusCode: http.StatusOK,
}
var r500 = http.Response{
	Status:     http.StatusText(http.StatusInternalServerError),
	StatusCode: http.StatusInternalServerError,
}

var p = []*entities.Pokemon{
	{
		Name: "test1",
		ID:   1,
	},
	{
		Name: "test2",
		ID:   2,
	},
	{
		Name: "test3",
		ID:   3,
	},
}

type mockGetPokemonUseCase struct {
	mock.Mock
}

func (m mockGetPokemonUseCase) GetPokemon(id string) (*entities.Pokemon, error) {
	args := m.Called(id)
	return args[0].(*entities.Pokemon), args.Error(1)
}

type mockGetPokemonsUseCase struct {
	mock.Mock
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
			err:          errors.New("testing"),
			url:          "/pokemons",
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
			data, _ := ioutil.ReadAll(res.Body)
			pjson, _ := json.Marshal(p)
			if tc.hasError {
				assert.EqualValues(t, tc.httpResponse.StatusCode, res.StatusCode)
			} else {
				assert.EqualValues(t, (string(pjson) + "\n"), string(data))
				assert.EqualValues(t, tc.httpResponse.StatusCode, res.StatusCode)
			}
		})
	}
}

func TestGetPokemon(t *testing.T) {
	testCases := []struct {
		name         string
		response     *entities.Pokemon
		httpResponse http.Response
		hasError     bool
		err          error
		url          string
		id           string
	}{
		{
			name:         "No error",
			response:     p[0],
			httpResponse: r200,
			hasError:     false,
			err:          nil,
			url:          "/pokemons/1",
			id:           "1",
		},
		{
			name:         "error",
			response:     nil,
			httpResponse: r500,
			hasError:     true,
			err:          errors.New("Testing"),
			url:          "/pokemons",
			id:           "1",
		},
		{
			name:         "letter id parameter",
			response:     nil,
			httpResponse: r400,
			hasError:     false,
			err:          nil,
			url:          "/pokemons/a",
			id:           "a",
		},
		{
			name:         "id doesn't exist in csv",
			response:     p[0],
			httpResponse: r200,
			hasError:     false,
			err:          nil,
			url:          "/pokemons/5",
			id:           "5",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mgp := mockGetPokemonsUseCase{}
			mgps := mockGetPokemonUseCase{}
			mgps.On("GetPokemon", tc.id).Return(tc.response, tc.err)
			vars := map[string]string{
				"id": tc.id,
			}
			h := New(mgp, mgps)
			r := httptest.NewRequest(http.MethodGet, tc.url, nil)
			r = mux.SetURLVars(r, vars)
			w := httptest.NewRecorder()
			h.GetPokemon(w, r)
			res := w.Result()
			data, _ := ioutil.ReadAll(res.Body)
			pjson, _ := json.Marshal(tc.response)
			switch res.StatusCode {
			case http.StatusBadRequest:
				msgbyte, _ := json.Marshal(map[string]string{"message": "Bad request: ID is not valid"})
				assert.EqualValues(t, tc.httpResponse.StatusCode, res.StatusCode)
				assert.EqualValues(t, string(msgbyte)+"\n", string(data))
			case http.StatusInternalServerError:
				msgbyte, _ := json.Marshal(map[string]string{"message": "Ooops!! :(", "error": tc.err.Error()})
				assert.EqualValues(t, tc.httpResponse.StatusCode, res.StatusCode)
				assert.EqualValues(t, string(msgbyte)+"\n", string(data))
			case http.StatusNotFound:
				if tc.hasError {
					assert.EqualValues(t, map[string]string{"message": "Pokemons Not Found", "error": "Not found"}, res.StatusCode)
					assert.EqualValues(t, (string(pjson) + "\n"), string(data))
				} else {
					assert.EqualValues(t, map[string]string{"message": "Pokemons Not Found", "error": tc.err.Error()}, res.StatusCode)
					assert.EqualValues(t, (string(pjson) + "\n"), string(data))
				}
			case http.StatusOK:
				assert.EqualValues(t, tc.httpResponse.StatusCode, res.StatusCode)
				assert.EqualValues(t, (string(pjson) + "\n"), string(data))
			}
			if tc.hasError {
				assert.EqualValues(t, tc.httpResponse.StatusCode, res.StatusCode)
			}
		})
	}
}
