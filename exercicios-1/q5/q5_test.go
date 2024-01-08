package q5

import (
	"testing"
)

func TestProcessString(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "CEUB", want: ".c.b"},
		{input: "ceub", want: ".c.b"},
		{input: "aBAcAba", want: ".b.c.b"},
		{input: "Codeforces", want: ".c.d.f.r.c.s"},
		{input: "codeforces", want: ".c.d.f.r.c.s"},
		{input: "a", want: ""},
		{input: "A", want: ""},
		{input: "tour", want: ".t.r"},
		{input: "TOUR", want: ".t.r"},
		{input: "tourist", want: ".t.r.s.t"},
		{input: "TOURIST", want: ".t.r.s.t"},
		{input: "tourists", want: ".t.r.s.t.s"},
		{input: "AISDJasikdjnasodnhAIOOS", want: ".s.d.j.s.k.d.j.n.s.d.n.h.s"},
		{input: "AJISASJIjasiasjAIsjASJIAIAJSAJISASNMDKLEQWIOfNAKSJLDnmQOWELASKDMNQOIEWFNLAAnajksnfoaiwsLANK", want: ".j.s.s.j.j.s.s.j.s.j.s.j.j.s.j.s.s.n.m.d.k.l.q.w.f.n.k.s.j.l.d.n.m.q.w.l.s.k.d.m.n.q.w.f.n.l.n.j.k.s.n.f.w.s.l.n.k"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := ProcessString(test.input)
			if got != test.want {
				t.Errorf("ProcessString(%q) = %q, want %q", test.input, got, test.want)
			}
		})
	}
}
