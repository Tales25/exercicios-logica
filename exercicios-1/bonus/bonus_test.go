package bonus

import (
	"fmt"
	"testing"
)

func TestCalculateTowers(t *testing.T) {
	tests := []struct {
		input     []int
		maxHeight int
		towers    int
	}{
		{input: []int{1, 2, 3, 4, 5}, maxHeight: 1, towers: 5},
		{input: []int{5, 6, 7, 6}, maxHeight: 2, towers: 3},
		{input: []int{1, 2, 3, 3}, maxHeight: 2, towers: 3},
		{input: []int{1, 2, 3, 3, 3}, maxHeight: 3, towers: 3},
		{input: []int{3, 1, 2, 1}, maxHeight: 2, towers: 3},
		{input: []int{1}, maxHeight: 1, towers: 1},
		{input: []int{1, 1000, 1000, 1000}, maxHeight: 3, towers: 2},
		{input: []int{1000, 1000, 1000, 8, 7}, maxHeight: 3, towers: 3},
		{input: []int{5, 5, 5, 5, 5}, maxHeight: 5, towers: 1},
		{input: []int{112, 277, 170, 247, 252, 115, 157, 293, 256, 143, 196, 90, 12, 164, 164, 42, 8, 223, 167, 109, 175, 232, 239, 111, 148, 51, 9, 254, 93, 32, 268, 162, 231, 91, 47, 162, 161, 191, 195, 145, 247, 292, 129, 199, 230, 94, 144, 217, 18, 205, 176, 20, 143, 198, 121, 243, 211, 262, 230, 277, 195, 255, 108, 290, 220, 275, 158, 2, 286, 200, 60, 267, 278, 207, 123, 150, 123, 116, 131, 13, 12, 226, 33, 244, 30, 275, 263, 45, 158, 192, 254, 149, 242, 176, 62, 224, 221, 288, 250, 160, 155, 225, 132, 143, 276, 293, 218, 145, 197, 175, 33, 129, 79, 206, 210, 192, 222, 262, 190, 52, 274, 243, 233}, maxHeight: 3, towers: 101},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("input=%v", test.input), func(t *testing.T) {
			maxHeight, towers := CalculateTowers(test.input)
			if maxHeight != test.maxHeight {
				t.Errorf("Altura máxima esperada: %d, Altura máxima obtida: %d", test.maxHeight, maxHeight)
			}
			if towers != test.towers {
				t.Errorf("Torres esperadas: %d, Torres obtidas: %d", test.towers, towers)
			}
		})
	}
}
