package q3

import (
	"fmt"
	"testing"
)

func TestDominoPieces(t *testing.T) {
	tests := []struct {
		m       int
		n       int
		want    int
		wantErr bool
	}{
		{m: 0, n: 0, want: 0, wantErr: true},
		{m: -1, n: 1, want: 0, wantErr: true},
		{m: 1, n: -1, want: 0, wantErr: true},
		{m: 2, n: 4, want: 4, wantErr: false},
		{m: 3, n: 3, want: 4, wantErr: false},
		{m: 1, n: 5, want: 2, wantErr: false},
		{m: 1, n: 6, want: 3, wantErr: false},
		{m: 1, n: 15, want: 7, wantErr: false},
		{m: 1, n: 16, want: 8, wantErr: false},
		{m: 2, n: 5, want: 5, wantErr: false},
		{m: 15, n: 15, want: 112, wantErr: false},
		{m: 14, n: 16, want: 112, wantErr: false},
		{m: 11, n: 13, want: 71, wantErr: false},
		{m: 1, n: 1, want: 0, wantErr: false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("m=%d,n=%d", test.m, test.n), func(t *testing.T) {
			got, err := DominoPieces(test.m, test.n)
			if (err != nil) != test.wantErr {
				t.Errorf("DominoPieces() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("DominoPieces() got = %v, want %v", got, test.want)
			}
		})
	}
}
