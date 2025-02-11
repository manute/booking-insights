package stats

import (
	"testing"
)

func Test_ProfitsPerNight(t *testing.T) {
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
			name: "given zero inputs",
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
			in:   []float64{10, 12.1, 10.29},
			want: profits{avg: 10.796666666666667, min: 10, max: 12.1},
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
