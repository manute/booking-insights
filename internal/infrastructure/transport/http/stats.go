package http

import (
	"booking-insights/internal/domain/stats"
	"encoding/json"
	"net/http"
)

type ServiceStats interface {
	ProfitsPerNight(payload []stats.ProfitsPerNightReqDTO) stats.ProfitsPerNightRespDTO
}

type StatsHandler struct {
	service ServiceStats
}

func NewStatsHandler(svc ServiceStats) *StatsHandler {
	return &StatsHandler{service: svc}
}

func (h *StatsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		e := errorResp{"method not accepted, only POST"}
		JSONError(w, &e, http.StatusMethodNotAllowed)
		return
	}
	h.postStats(w, r)
	return
}

func (h *StatsHandler) postStats(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var payload []statsPayload
	if err := decoder.Decode(&payload); err != nil {
		e := errorResp{"decoding json payload: " + err.Error()}
		JSONError(w, &e, http.StatusMethodNotAllowed)
		return
	}

	data := mapperReq(payload)
	profits := h.service.ProfitsPerNight(data)

	res := mapperRes(profits)
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(&res); err != nil {
		e := errorResp{"encoding json resp: " + err.Error()}
		JSONError(w, &e, http.StatusBadRequest)
		return
	}
}

func mapperReq(pp []statsPayload) []stats.ProfitsPerNightReqDTO {
	var out []stats.ProfitsPerNightReqDTO
	for _, p := range pp {
		out = append(out, stats.ProfitsPerNightReqDTO{
			SellingRate: float64(p.SellingRate),
			Margin:      p.Margin,
			Nights:      p.Nights,
		})
	}
	return out
}

func mapperRes(dto stats.ProfitsPerNightRespDTO) statsResp {
	return statsResp{
		Max: dto.Max,
		Min: dto.Min,
		Avg: dto.Avg,
	}
}
