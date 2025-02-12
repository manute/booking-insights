package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errorResp struct {
	Message string `json:"message"`
}

type statsPayload struct {
	RequestID   string `json:"request_id"`
	CheckIn     string `json:"check_in"`
	Nights      int    `json:"nights"`
	SellingRate int    `json:"selling_rate"`
	Margin      int    `json:"margin"`
}

type statsResp struct {
	Max float64 `json:"max_night"`
	Min float64 `json:"min_night"`
	Avg float64 `json:"avg_night"`
}

type MaximizeResp struct {
	ReqIDs      []string `json:"request_ids"`
	TotalProfit float64  `json:"total_profit"`
	Max         float64  `json:"max_night"`
	Min         float64  `json:"min_night"`
	Avg         float64  `json:"avg_night"`
}

func newErrorResp(err error) string {
	return fmt.Sprintf("{\"message\":\"%s\"}", err.Error())
}

// JSONError makes possible the json as a response
// because the http.Error overwrites the header value
func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
