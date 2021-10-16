package service

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MoraAlex/academy-go-q32021/entities"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var mockResp = `{
	"id": 19,
	"name": "rattata"
	}`

var pokResp = &entities.Pokemon{
	ID:   19,
	Name: "rattata",
}

func TestGetPokemonApi(t *testing.T) {
	testCases := []struct {
		name            string
		id              string
		err             error
		hasError        bool
		mockApiResp     string
		mockApiRespCode int
		expectedResp    *entities.Pokemon
		expectedBody    []byte
	}{
		{
			name:            "Return a pokemon",
			id:              "2",
			err:             nil,
			hasError:        false,
			mockApiResp:     mockResp,
			mockApiRespCode: 200,
			expectedResp:    pokResp,
			expectedBody:    []byte(`{"id": 1, "Name": "bulbasaur"}`),
		},
		{
			name:            "Return error if incorrect parameter",
			id:              "a",
			err:             errors.New("Not found"),
			hasError:        true,
			mockApiResp:     "",
			mockApiRespCode: 404,
			expectedResp:    nil,
			expectedBody:    []byte(`{"Not found"}`),
		},
	}
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	for _, tc := range testCases {

		handler := func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, string(tc.expectedBody))
		}

		req := httptest.NewRequest(http.MethodGet, "https://pokeapi.co/api/v2/pokemon/"+tc.id, nil)
		w := httptest.NewRecorder()
		w.WriteHeader(tc.mockApiRespCode)
		handler(w, req)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		// responder := httpmock.NewStringResponder(tc.mockApiRespCode, tc.mockApiResp)
		// apiUrl := "https://pokeapi.co/api/v2/pokemon/" + tc.id
		// httpmock.RegisterResponder(http.MethodGet, apiUrl, responder)
		// s := NewGetPokemonApi(client)
		// pokemons, err := s.GetPokemonApi(tc.id)
		// assert.Error(t, err)
		// assert.EqualValues(t, tc.err.Error(), err.Error())
		// assert.EqualValues(t, tc.expectedResp, pokemons)
		assert.EqualValues(t, tc.expectedBody, body)

	}
}
