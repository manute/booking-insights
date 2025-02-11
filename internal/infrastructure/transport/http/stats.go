package http

import (
	"booking-insights/internal/domain/stats"
	"encoding/json"
	"net/http"
)

type ServiceStats interface {
	ProfitsPerNight(xs []float64) stats.ProfitsDTO
}

type StatsHandler struct {
	service stats.Service
}

func NewStatsHandler(svc stats.Service) *StatsHandler {
	return &StatsHandler{service: svc}
}

func (h *StatsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	h.post(w, r)
	return
}

func (h *StatsHandler) post(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var payload []statsPayload
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, "decoding json payload", http.StatusBadRequest)
		return
	}

	// TODO: this is domain?
	var data []float64
	for _, p := range payload {
		d := (p.SellingRate * p.Margin / 100) / p.Nights
		data = append(data, float64(d))
	}

	profits := h.service.ProfitsPerNight(data)
	res := mapper(profits)

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(&res); err != nil {
		http.Error(w, "encoding json resp", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

type ProfitsResp struct {
	Max float64 `json:"max"`
	Min float64 `json:"min"`
	Avg float64 `json:"avg"`
}

func mapper(dto stats.ProfitsDTO) ProfitsResp {
	return ProfitsResp{
		Max: dto.Max,
		Min: dto.Min,
		Avg: dto.Avg,
	}
}
