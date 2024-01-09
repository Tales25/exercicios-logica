package q4

import (
	"testing"
)

func TestDestino(t *testing.T) {
	tests := []struct {
		name     string
		caminhos [][2]string
		want     string
		wantErr  bool
	}{
		{
			name:     "empty paths",
			caminhos: [][2]string{},
			want:     "",
			wantErr:  true,
		},
		{
			name: "one path",
			caminhos: [][2]string{
				{"A", "B"},
			},
			want: "B",
		},
		{
			name: "two paths",
			caminhos: [][2]string{
				{"A", "B"},
				{"B", "C"},
			},
			want: "C",
		},
		{
			name: "three paths",
			caminhos: [][2]string{
				{"C", "D"},
				{"A", "B"},
				{"B", "C"},
			},
			want: "D",
		},
		{
			name: "four paths",
			caminhos: [][2]string{
				{"A", "B"},
				{"D", "E"},
				{"B", "C"},
				{"C", "D"},
			},
			want: "E",
		},
		{
			name: "real path",
			caminhos: [][2]string{
				{"London", "New York"},
				{"New York", "Lima"},
				{"Lima", "Sao Paulo"},
			},
			want: "Sao Paulo",
		},
		{
			name: "no destination",
			caminhos: [][2]string{
				{"A", "B"},
				{"B", "C"},
				{"C", "A"},
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			destino, err := Destino(test.caminhos)
			if test.wantErr && err == nil {
				t.Errorf("Erro esperado, mas nenhum erro foi retornado")
			}

			if destino != test.want {
				t.Errorf("Destino() = %v, want %v", destino, test.want)
			}
		})
	}
}
