package http_test

import (
	"booking-insights/internal/domain/stats"
	httptransport "booking-insights/internal/infrastructure/transport/http"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"sort"
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

func Test_MaximizeHandler(t *testing.T) {
	tests := []struct {
		name     string
		inReq    string
		wantResp string
	}{
		{"happy path", "maximize.json", "maximize.golden"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in, err := os.ReadFile(filepath.Join("testdata", tt.inReq))
			if err != nil {
				t.Errorf("expected error to be nil got %s", err)
			}

			payload := bytes.NewBuffer(in)
			req := httptest.NewRequest(http.MethodPost, "/maximize", payload)

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

			dg := json.NewDecoder(bytes.NewBuffer(respData))
			var got httptransport.MaximizeResp
			if err = dg.Decode(&got); err != nil {
				t.Errorf("expected error to be nil got %s", err)
			}

			dw := json.NewDecoder(bytes.NewBuffer(respWant))
			var want httptransport.MaximizeResp
			if err = dw.Decode(&want); err != nil {
				t.Errorf("expected error to be nil got %s", err)
			}

			// for strigns slices comnparision
			sort.Strings(got.ReqIDs)
			sort.Strings(want.ReqIDs)

			if !reflect.DeepEqual(got.ReqIDs, want.ReqIDs) {
				t.Errorf("expected a %v, instead got: %v", want.ReqIDs, got.ReqIDs)
			}

		})
	}
}
