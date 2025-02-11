package stats

import (
	"testing"
)

func Test_Profits(t *testing.T) {

	tests := []struct {
		name string
		in   []float64
		want profits
	}{
		{
			name: "given nil input",
			want: profits{0, 0, 0},
		},
		{
			name: "given emptys",
			in:   []float64{0},
			want: profits{0, 0, 0},
		},
		{
			name: "given two profits",
			in:   []float64{8, 8.58},
			want: profits{avg: 8.29, min: 8, max: 8.58},
		},
		{
			name: "given 3 profits",
			in:   []float64{10, 12.1, 10.80},
			want: profits{avg: 10.966666666666669, min: 10.0, max: 12.10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := newProfit(tt.in)
			if want, got := tt.want, p; want != got {
				t.Errorf("expected a %f, instead got: %f", want, got)
			}

		})
	}
}
