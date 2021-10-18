package repository

import (
	"encoding/csv"
	"errors"
	"os"
	"testing"

	"github.com/MoraAlex/academy-go-q32021/entities"
	"github.com/stretchr/testify/assert"
)

var p = []*entities.Pokemon{
	{
		ID:   1,
		Name: "Test1",
	},
	{
		ID:   2,
		Name: "Test2",
	},
	{
		ID:   3,
		Name: "Test3",
	},
}

func TestGetAll(t *testing.T) {
	testCases := []struct {
		name     string
		response []*entities.Pokemon
		length   int
		err      error
		hasError bool
		testFile string
	}{
		{
			name:     "Get all pokemons",
			response: p,
			length:   1,
			err:      nil,
			hasError: false,
			testFile: "./testfile.csv",
		},
		{
			name:     "Error",
			response: nil,
			err:      errors.New("Test"),
			hasError: true,
			testFile: "",
		},
	}

	for _, tc := range testCases {
		if _, err := os.Stat("./testfile.csv"); os.IsNotExist(err) {
			f, err := os.Create("testfile.csv")
			if err != nil {
				panic(err)
			}
			// Escribir BOM UTF-8
			f.WriteString("\xEF\xBB\xBF")
			// Crea una nueva secuencia de archivos de escritura
			w := csv.NewWriter(f)
			data := [][]string{
				{"id", "Name", "Type 1", "Type 2"},
				{"1", "Test1", "", ""},
				{"2", "Test2", "", ""},
				{"3", "Test3", "", ""},
			}
			//Entrada de datos
			w.WriteAll(data)
			w.Flush()
			f.Close()
		}
		t.Run(tc.name, func(t *testing.T) {
			s := NewPokemon(tc.testFile)
			pokemons, err := s.GetAll()
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.EqualValues(t, p, pokemons)
			}

		})
	}
	os.Remove("./testfile.csv")
}
