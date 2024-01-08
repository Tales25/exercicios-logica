package q1

import (
	"fmt"
	"testing"
)

func TestDivideWatermelon(t *testing.T) {
	tests := []struct {
		weight      int
		expected    bool
		expectedErr bool
	}{
		{weight: -1, expected: false, expectedErr: true},
		{weight: 0, expected: false, expectedErr: true},
		{weight: 1, expected: false, expectedErr: false},
		{weight: 2, expected: false, expectedErr: false},
		{weight: 3, expected: false, expectedErr: false},
		{weight: 4, expected: true, expectedErr: false},
		{weight: 5, expected: false, expectedErr: false},
		{weight: 6, expected: true, expectedErr: false},
		{weight: 7, expected: false, expectedErr: false},
		{weight: 8, expected: true, expectedErr: false},
		{weight: 9, expected: false, expectedErr: false},
		{weight: 10, expected: true, expectedErr: false},
		{weight: 11, expected: false, expectedErr: false},
		{weight: 12, expected: true, expectedErr: false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("weight=%d", test.weight), func(t *testing.T) {
			result, err := DivideWatermelon(test.weight)
			if test.expectedErr && err == nil {
				t.Errorf("Erro esperado, mas nenhum erro foi retornado")
			}

			if result != test.expected {
				t.Errorf("Resultado esperado: %t, Resultado obtido: %t", test.expected, result)
			}
		})
	}
}
