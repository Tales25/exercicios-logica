package q4

import (
	"fmt"
	"testing"
)

func TestClassifyPrices(t *testing.T) {
	tests := []struct {
		prices  []int
		want    int
		wantErr bool
	}{
		{prices: []int{1, 2, 3, 4, 5}, want: 1, wantErr: false},
		{prices: []int{5, 4, 3, 2, 1}, want: 2, wantErr: false},
		{prices: []int{1, 2, 3, 4, 3}, want: 3, wantErr: false},
		{prices: []int{}, want: 0, wantErr: true},
		{prices: []int{1}, want: 3, wantErr: false},
		{prices: []int{1, 20, 60, 90}, want: 1, wantErr: false},
		{prices: []int{10, 5, 2}, want: 2, wantErr: false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("prices=%v", test.prices), func(t *testing.T) {
			got, err := ClassifyPrices(test.prices)
			if (err != nil) != test.wantErr {
				t.Errorf("ClassifyPrices() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("ClassifyPrices() got = %v, want %v", got, test.want)
			}
		})
	}
}
