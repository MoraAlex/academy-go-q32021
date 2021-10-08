package service

import (
	"errors"
	"testing"

	"github.com/MoraAlex/academy-go-q32021/entities"
	"github.com/stretchr/testify/assert"
)

var p = entities.Pokemon{
	ID:   1,
	Name: "Test1",
}

func TestUpdateCsv(t *testing.T) {
	testCases := []struct {
		name     string
		response error
		hasError bool
	}{
		{
			name:     "save the pokemon",
			response: nil,
			hasError: false,
		},
		{
			name:     "Doesn't save the same pokemon twice",
			response: errors.New("test"),
			hasError: false,
		},
		{
			name:     "Doesn't overwrite the file",
			response: nil,
			hasError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewUpdateCsv()
			pokemons, err := s.UpdateCsv(p)
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.EqualValues(t, p, pokemons)
			}

		})

	}
}
