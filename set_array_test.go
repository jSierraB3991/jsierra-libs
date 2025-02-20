package eliotlibs

import (
	"reflect"
	"testing"
)

// Test para verificar que la función agrega elementos sin duplicados
func TestAddUniqueNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    []uint
		toAdd    uint
		expected []uint
	}{
		{"Agregar número nuevo", []uint{1, 2, 3}, 4, []uint{1, 2, 3, 4}},
		{"Agregar número existente", []uint{1, 2, 3}, 2, []uint{1, 2, 3}},
		{"Agregar a lista vacía", []uint{}, 5, []uint{5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddUniqueUNumber(tt.input, tt.toAdd)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Esperado %v, pero obtuvo %v", tt.expected, result)
			}
		})
	}
}

func TestAddUniqueString(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		toAdd    string
		expected []string
	}{
		{"Agregar string nuevo", []string{"a", "b", "c"}, "d", []string{"a", "b", "c", "d"}},
		{"Agregar string existente", []string{"a", "b", "c"}, "b", []string{"a", "b", "c"}}, // No cambia
		{"Agregar a lista vacía", []string{}, "z", []string{"z"}},                           // Agrega primer string
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddUniqueString(tt.input, tt.toAdd)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Esperado %v, pero obtuvo %v", tt.expected, result)
			}
		})
	}
}
