package q2

import (
	"fmt"
	"testing"
)

func TestCalculateTeams(t *testing.T) {
	tests := []struct {
		mathematicians int
		programmers    int
		want           int
	}{
		{
			mathematicians: 5,
			programmers:    5,
			want:           2,
		},
		{
			mathematicians: 10,
			programmers:    1,
			want:           1,
		},
		{
			mathematicians: 2,
			programmers:    3,
			want:           1,
		},
		{
			mathematicians: 0,
			programmers:    0,
			want:           0,
		},
		{
			mathematicians: 17,
			programmers:    2,
			want:           2,
		},
		{
			mathematicians: 10000,
			programmers:    10000,
			want:           5000,
		},
		{
			mathematicians: 1000,
			programmers:    0,
			want:           0,
		},
		{
			mathematicians: 0,
			programmers:    1000,
			want:           0,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("mathematicians=%d,programmers=%d", test.mathematicians, test.programmers), func(t *testing.T) {
			participants := make([]Participant, test.mathematicians+test.programmers)
			var i int
			for i = 0; i < test.mathematicians; i++ {
				participants[i] = createParticipant("Mathematician")
			}
			for i = test.mathematicians; i < test.mathematicians+test.programmers; i++ {
				participants[i] = createParticipant("Programmer")
			}

			got := CalculateTeams(participants)
			if got != test.want {
				t.Errorf("ProblemsSolved() got = %v, want %v", got, test.want)
			}
		})
	}
}

func createParticipant(role string) Participant {
	return Participant{
		Role: role,
	}
}
