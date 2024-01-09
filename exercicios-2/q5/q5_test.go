package q5

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRegister(t *testing.T) {
	tests := []struct {
		names []string
		want  []string
	}{
		{
			names: []string{"João", "Maria", "José"},
			want:  []string{"OK", "OK", "OK"},
		},
		{
			names: []string{"João", "Maria", "José", "João"},
			want:  []string{"OK", "OK", "OK", "João1"},
		},
		{
			names: []string{"abacaba", "acaba", "abacaba", "acab"},
			want:  []string{"OK", "OK", "abacaba1", "OK"},
		},
		{
			names: []string{"abacaba", "acaba", "abacaba", "acab", "abacaba"},
			want:  []string{"OK", "OK", "abacaba1", "OK", "abacaba2"},
		},
		{
			names: []string{"first", "second", "third", "second", "first", "second", "third", "second", "first"},
			want:  []string{"OK", "OK", "OK", "second1", "first1", "second2", "third1", "second3", "first2"},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.names), func(t *testing.T) {
			got := Register(test.names)
			if len(test.want) == 0 && len(got) == 0 {
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Register(%v) = %v, want %v", test.names, got, test.want)
			}
		})
	}
}
