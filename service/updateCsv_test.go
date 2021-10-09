package service

import (
	"encoding/csv"
	"errors"
	"os"
	"testing"

	"github.com/MoraAlex/academy-go-q32021/entities"
	"github.com/stretchr/testify/assert"
)

var p = []*entities.Pokemon{{
	ID:   1,
	Name: "Test1",
}}

func TestUpdateCsv(t *testing.T) {
	testCases := []struct {
		name     string
		response []*entities.Pokemon
		length   int
		err      error
		hasError bool
	}{
		{
			name:     "save the pokemon",
			response: p,
			length:   1,
			err:      nil,
			hasError: false,
		},
		{
			name:     "Doesn't save the same pokemon twice",
			response: nil,
			err:      errors.New("Test"),
			hasError: true,
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
			}
			//Entrada de datos
			w.WriteAll(data)
			w.Flush()
			f.Close()
		}
		t.Run(tc.name, func(t *testing.T) {
			s := NewUpdateCsv("./testfile.csv")
			pokemons, err := s.UpdateCsv(*p[0])
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.EqualValues(t, p, pokemons)
			}

		})
	}
	os.Remove("./testfile.csv")
}
