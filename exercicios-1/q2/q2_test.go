package q2

import (
	"fmt"
	"testing"
)

func TestProblemsSolved(t *testing.T) {
	tests := []struct {
		answers [][3]bool
		want    int
	}{
		{
			answers: [][3]bool{
				{true, true, true},
				{true, true, false},
				{true, false, false},
			},
			want: 2,
		},
		{
			answers: [][3]bool{
				{false, true, false},
				{false, true, true},
				{false, false, false},
			},
			want: 1,
		},
		{
			answers: [][3]bool{
				{false, true, false},
				{false, true, true},
				{false, false, false},
				{true, true, true},
				{true, true, false},
			},
			want: 3,
		},
		{
			answers: [][3]bool{
				{false, true, false},
			},
			want: 0,
		},
		{
			answers: [][3]bool{
				{true, false, false},
				{true, true, true},
			},
			want: 1,
		},
		{
			answers: [][3]bool{
				{false, false, false},
				{false, true, true},
				{true, true, true},
				{false, true, false},
				{true, false, true},
				{true, true, true},
				{false, false, true},
				{true, false, false},
				{true, true, false},
				{true, false, true},
				{false, true, false},
				{false, false, true},
				{true, true, false},
				{false, true, false},
				{true, true, false},
				{false, false, false},
				{true, true, true},
				{true, false, true},
				{false, false, true},
				{true, true, false},
				{true, true, true},
				{false, true, true},
				{true, true, false},
				{false, false, false},
				{false, false, false},
				{true, true, true},
				{false, false, false},
				{true, true, true},
				{false, true, true},
				{false, false, true},
				{false, false, false},
				{false, false, false},
				{true, true, false},
				{true, true, false},
				{true, false, true},
				{true, false, false},
				{true, false, true},
				{true, false, true},
				{false, true, true},
				{true, true, false},
				{true, true, false},
				{false, true, false},
				{true, false, true},
				{false, false, false},
				{false, false, false},
				{false, false, false},
				{false, false, true},
				{true, true, true},
				{false, true, true},
				{true, false, true},
			},
			want: 29,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			if got := ProblemsSolved(test.answers); got != test.want {
				t.Errorf("ProblemsSolved() = %v, want %v", got, test.want)
			}
		})
	}
}
