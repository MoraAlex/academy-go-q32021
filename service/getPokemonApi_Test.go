package service

// import (
// 	"testing"

// 	"github.com/MoraAlex/academy-go-q32021/entities"
// )

// var p = entities.Pokemon{
// 	Name:       "Bulbasaur",
// 	ID:         1,
// 	MainType:   "grass",
// 	SecondType: "Poison",
// }

// func TestgetPokemonApiService(t *testing.T) {
// 	testCases := []struct {
// 		name           string
// 		expectedLength int
// 		response       entities.Pokemon
// 		hasError       bool
// 		error          error
// 		id             string
// 	}{
// 		{
// 			name:           "return pokemon if id exist",
// 			expectedLength: 1,
// 			response:       p,
// 			hasError:       false,
// 			error:          nil,
// 			id:             "1",
// 		},
// 		{
// 			name:           "Not found if incorrect param",
// 			expectedLength: 1,
// 			response:       p,
// 			hasError:       false,
// 			error:          nil,
// 			id:             "a",
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			mock :=
// 		})
// 	}
// }
