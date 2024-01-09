package bonus

import (
	"fmt"
	"testing"
)

func TestCalculateAveragePrice(t *testing.T) {
	tests := []struct {
		shirts  []Shirt
		wantMax float64
		wantMin float64
		wantErr bool
	}{
		{
			shirts: []Shirt{
				createShirt("M", 10.0),
				createShirt("M", 20.0),
				createShirt("M", 30.0),
				createShirt("XXXXXS", 30.0),
				createShirt("XXXXXS", 30.0),
				createShirt("XXXXXS", 30.0),
				createShirt("XXS", 350.0),
				createShirt("XXXS", 450.0),
				createShirt("S", 150.0),
				createShirt("S", 550.0),
			},
			wantMax: 20.0,
			wantMin: 30.0,
		},
		{
			shirts: []Shirt{
				createShirt("M", 10.0),
				createShirt("M", 20.0),
				createShirt("M", 30.0),
				createShirt("XXXS", 30.0),
				createShirt("XXS", 30.0),
				createShirt("S", 30.0),
				createShirt("L", 40.0),
				createShirt("L", 50.0),
				createShirt("L", 60.0),
			},
			wantMax: 50.0,
			wantMin: 30.0,
		},
		{
			shirts: []Shirt{
				createShirt("XXXXXXXXXS", 20.0),
				createShirt("XXS", 30.0),
				createShirt("XXXXXXXXXS", 30.0),
				createShirt("XXXXL", 40.0),
				createShirt("XXXXL", 50.0),
				createShirt("XXL", 60.0),
				createShirt("M", 160.0),
				createShirt("M", 60.0),
			},
			wantMax: 45.0,
			wantMin: 25.0,
		},
		{
			shirts: []Shirt{
				createShirt("XXXXL", 40.0),
				createShirt("XXXXXXL", 50.0),
				createShirt("XXL", 30.0),
				createShirt("XXL", 20.0),
				createShirt("XXL", 25.0),
				createShirt("XXXXXXL", 55.0),
				createShirt("XXXXXXL", 60.0),
				createShirt("XXXXL", 60.0),
			},
			wantMax: 55.0,
			wantMin: 25.0,
		},
		{
			shirts: []Shirt{
				createShirt("XXXXS", 40.0),
				createShirt("XXXXXXS", 50.0),
				createShirt("XXS", 30.0),
				createShirt("XXS", 20.0),
				createShirt("XXS", 25.0),
				createShirt("XXXXXXS", 55.0),
				createShirt("XXXXXXS", 60.0),
				createShirt("XXXXS", 60.0),
			},
			wantMax: 25.0,
			wantMin: 55.0,
		},
		{
			shirts:  []Shirt{},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("shirts=%v", test.shirts), func(t *testing.T) {
			max, min, err := CalculateAveragePrice(test.shirts)
			if (err != nil) != test.wantErr {
				t.Errorf("CalculateAveragePrice() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if max != test.wantMax {
				t.Errorf("CalculateAveragePrice() max = %v, wantMax %v", max, test.wantMax)
			}
			if min != test.wantMin {
				t.Errorf("CalculateAveragePrice() min = %v, wantMin %v", min, test.wantMin)
			}
		})
	}
}

func createShirt(size string, price float64) Shirt {
	return Shirt{
		Size:  size,
		Price: price,
	}
}
