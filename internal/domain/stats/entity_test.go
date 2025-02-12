package stats

import (
	"reflect"
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

func Test_Maximize(t *testing.T) {
	in := []MaximizeReqDTO{
		{ReqID: "xyz", CheckIn: "2020-01-01", Nights: 5, SellingRate: 200, Margin: 20},
		{ReqID: "yja", CheckIn: "2020-01-10", Nights: 4, SellingRate: 160, Margin: 30},
	}

	got, err := maximize(in)
	if err != nil {
		t.Error("not expected error", err)
	}

	want := &MaximizeRespDTO{
		ReqsID:      []string{"xyz", "yja"},
		TotalProfit: 88,
		Max:         12,
		Min:         8,
		Avg:         10,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected a %v, instead got: %v", want, got)
	}

}
