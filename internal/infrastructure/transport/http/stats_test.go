package http_test

import (
	"booking-insights/internal/domain/stats"
	httptransport "booking-insights/internal/infrastructure/transport/http"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_StatsHandlerHTTPMethods(t *testing.T) {

	tests := []struct {
		name         string
		inHttpMethod string
		wantHttpCode int
	}{
		{"GET", http.MethodGet, http.StatusMethodNotAllowed},
		{"PUT", http.MethodPut, http.StatusMethodNotAllowed},
		{"POST", http.MethodPost, http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			svcStats := stats.NewService()
			statsHandler := httptransport.NewStatsHandler(svcStats)

			req := httptest.NewRequest(tt.inHttpMethod, "/stats", bytes.NewBufferString("[]"))
			res := httptest.NewRecorder()
			statsHandler.ServeHTTP(res, req)
			fmt.Println("res", res.Body.String())
			if want, got := tt.wantHttpCode, res.Code; want != got {
				t.Errorf("expected a %d, instead got: %d", want, got)
			}
			res.Result().Body.Close()
		})
	}

}

func Test_StatsHandler(t *testing.T) {
	tests := []struct {
		name     string
		inReq    string
		wantResp string
	}{
		{"empty req", "stats_empty.json", "stats_empty.golden"},
		{"two bookings", "stats_two.json", "stats_two.golden"},
		{"three bookings", "stats_three.json", "stats_three.golden"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in, err := os.ReadFile(filepath.Join("testdata", tt.inReq))
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
			defer res.Result().Body.Close()

			respWant, err := os.ReadFile(filepath.Join("testdata", tt.wantResp))
			if err != nil {
				t.Errorf("expected error to be nil got %s", err)
			}

			g := strings.TrimRight(string(respData), "\n")
			w := strings.TrimRight(string(respWant), "\n")

			if want, got := w, g; strings.Compare(want, got) != 0 {
				t.Errorf("expected a %s, instead got: %s", want, got)
			}

		})
	}
}
