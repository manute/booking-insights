package stats

// import (
// 	// "booking-insights/internal/domain/stats"
// 	"testing"
// )

// func Test_Math(t *testing.T) {

// 	tests := []struct {
// 		name    string
// 		profits []float64
// 		wantAvg float64
// 	}{
// 		{
// 			name:    "given nil profits",
// 			profits: nil,
// 			wantAvg: 0,
// 		},
// 		{
// 			name:    "given emptyu profits",
// 			profits: []float64{},
// 			wantAvg: 0,
// 		},
// 		{
// 			name:    "given two profits",
// 			profits: []float64{8, 8.58},
// 			wantAvg: 8.29,
// 		},
// 		{
// 			name:    "given 3 profits",
// 			profits: []float64{10, 12.1, 10.80},
// 			wantAvg: 10.996667,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			avg := stats.Average(tt.profits)
// 			if want, got := tt.wantAvg, avg; want != got {
// 				t.Errorf("expected a %f, instead got: %f", want, got)
// 			}

// 		})
// 	}

// }
