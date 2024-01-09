package q1

import (
	"fmt"
	"testing"
)

func TestCanBuyFood(t *testing.T) {
	tests := []struct {
		stock      map[string]int
		dogs, cats int
		expected   bool
	}{
		{
			stock: map[string]int{
				"dog":       1,
				"cat":       1,
				"universal": 4,
			},
			dogs:     2,
			cats:     3,
			expected: true,
		},
		{
			stock: map[string]int{
				"dog":       0,
				"cat":       0,
				"universal": 0,
			},
			dogs:     0,
			cats:     0,
			expected: true,
		},
		{
			stock: map[string]int{
				"dog":       5,
				"cat":       5,
				"universal": 0,
			},
			dogs:     4,
			cats:     6,
			expected: false,
		},
		{
			stock: map[string]int{
				"dog":       50000000,
				"cat":       50000000,
				"universal": 100000000,
			},
			dogs:     100000000,
			cats:     100000000,
			expected: true,
		},
		{
			stock: map[string]int{
				"dog":       0,
				"cat":       0,
				"universal": 0,
			},
			dogs:     100000000,
			cats:     100000000,
			expected: false,
		},
		{
			stock: map[string]int{
				"dog":       1,
				"cat":       3,
				"universal": 2,
			},
			dogs:     2,
			cats:     5,
			expected: false,
		},
		{
			stock: map[string]int{
				"dog":       1,
				"cat":       3,
				"universal": 2,
			},
			dogs:     2,
			cats:     4,
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("available=%v,dogs=%d,cats=%d", test.stock, test.dogs, test.cats), func(t *testing.T) {
			got := CanBuyFood(test.stock, test.dogs, test.cats)
			if got != test.expected {
				t.Errorf("CanBuyFood() got = %v, want %v", got, test.expected)
			}
		})
	}
}
