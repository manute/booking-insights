package stats_test

import (
	"booking-insights/internal/domain/stats"
	"testing"
)

func Test_ServiceProfitsPerNight(t *testing.T) {
	tests := []struct {
		name string
		in   []stats.ProfitsPerNightReqDTO
		want stats.ProfitsPerNightRespDTO
	}{
		{
			name: "given nil input",
			want: stats.ProfitsPerNightRespDTO{0, 0, 0},
		},
		{
			name: "given zero inputs",
			in: []stats.ProfitsPerNightReqDTO{
				{SellingRate: 0, Margin: 0, Nights: 0},
			},
			want: stats.ProfitsPerNightRespDTO{0, 0, 0},
		},
		{
			name: "given two profits",
			in: []stats.ProfitsPerNightReqDTO{
				{SellingRate: 200, Margin: 20, Nights: 5},
				{SellingRate: 156, Margin: 22, Nights: 4},
			},
			want: stats.ProfitsPerNightRespDTO{Avg: 8.29, Min: 8, Max: 8.58},
		},
		{
			name: "given 3 profits",
			in: []stats.ProfitsPerNightReqDTO{
				{SellingRate: 50, Margin: 20, Nights: 1},
				{SellingRate: 55, Margin: 22, Nights: 1},
				{SellingRate: 49, Margin: 21, Nights: 1},
			},
			want: stats.ProfitsPerNightRespDTO{Max: 12.1, Min: 10, Avg: 10.796666666666667},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := stats.NewService()
			res := svc.ProfitsPerNight(tt.in)
			if want, got := tt.want, res; want != got {
				t.Errorf("expected a %f, instead got: %+v", want, got)
			}

		})
	}
}
