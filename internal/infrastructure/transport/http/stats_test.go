package http_test

import (
	"booking-insights/internal/domain/stats"
	httptransport "booking-insights/internal/infrastructure/transport/http"
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_StatsHandler(t *testing.T) {
	tests := []struct {
		name         string
		in           string
		want         string
		wantHttpCode int
	}{
		{"empty req", "empty_req.json", "empty_req.golden", http.StatusOK},
		{"two bookings", "bookings_req_2.json", "bookings_req_2.golden", http.StatusOK},
		{"three bookings", "bookings_req_3.json", "bookings_req_3.golden", http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in, err := os.ReadFile(filepath.Join("testdata", tt.in))
			if err != nil {
				t.Errorf("expected error to be nil got %s", err)
			}

			payload := bytes.NewBuffer(in)
			req := httptest.NewRequest(http.MethodPost, "/stats", payload)

			svcStats := stats.NewService()
			statsHandler := httptransport.NewStatsHandler(svcStats)

			res := httptest.NewRecorder()
			statsHandler.ServeHTTP(res, req)

			respData, err := io.ReadAll(res.Body)
			if err != nil {
				t.Errorf("expected error to be nil got %s", err)
			}

			respWant, err := os.ReadFile(filepath.Join("testdata", tt.want))
			if err != nil {
				t.Errorf("expected error to be nil got %s", err)
			}

			g := strings.TrimRight(string(respData), "\n")
			w := strings.TrimRight(string(respWant), "\n")

			if want, got := tt.wantHttpCode, res.Code; want != got {
				t.Errorf("expected a %d, instead got: %d", want, got)
			}
			if want, got := w, g; strings.Compare(want, got) != 0 {
				t.Errorf("expected a %s, instead got: %s", want, got)
			}

		})
	}
}
