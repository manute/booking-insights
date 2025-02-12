package http

import (
	"booking-insights/internal/domain/stats"
	"encoding/json"
	"net/http"
)

type ServiceStats interface {
	ProfitsPerNight(payload []stats.ProfitsPerNightReqDTO) stats.ProfitsPerNightRespDTO
	Maximize(payload []stats.MaximizeReqDTO) (stats.MaximizeRespDTO, error)
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

	path := r.URL.EscapedPath()

	if path == "/stats" {
		h.postStats(w, r)
	}
	if path == "/maximize" {
		h.postMaximize(w, r)
	}

	return
}

func (h *StatsHandler) postStats(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var payload []statsPayload
	if err := decoder.Decode(&payload); err != nil {
		e := errorResp{"decoding json payload: " + err.Error()}
		JSONError(w, &e, http.StatusBadRequest)
		return
	}

	data := mapperProfReq(payload)
	profits := h.service.ProfitsPerNight(data)

	res := mapperProfRes(profits)
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(&res); err != nil {
		e := errorResp{"encoding json resp: " + err.Error()}
		JSONError(w, &e, http.StatusInternalServerError)
		return
	}
}

func (h *StatsHandler) postMaximize(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var payload []statsPayload
	if err := decoder.Decode(&payload); err != nil {
		e := errorResp{"decoding json payload: " + err.Error()}
		JSONError(w, &e, http.StatusBadRequest)
		return
	}

	data := mapperMaxReq(payload)
	max, err := h.service.Maximize(data)
	if err != nil {
		e := errorResp{err.Error()}
		JSONError(w, &e, http.StatusInternalServerError)
		return
	}

	res := mapperMaxRes(max)
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(&res); err != nil {
		e := errorResp{"encoding json resp: " + err.Error()}
		JSONError(w, &e, http.StatusInternalServerError)
		return
	}

}

func mapperProfReq(pp []statsPayload) []stats.ProfitsPerNightReqDTO {
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

func mapperProfRes(dto stats.ProfitsPerNightRespDTO) statsResp {
	return statsResp{
		Max: dto.Max,
		Min: dto.Min,
		Avg: dto.Avg,
	}
}

func mapperMaxReq(pp []statsPayload) []stats.MaximizeReqDTO {
	var out []stats.MaximizeReqDTO
	for _, p := range pp {
		out = append(out, stats.MaximizeReqDTO{
			SellingRate: float64(p.SellingRate),
			Margin:      p.Margin,
			Nights:      p.Nights,
			ReqID:       p.RequestID,
			CheckIn:     p.CheckIn,
		})
	}
	return out
}

func mapperMaxRes(dto stats.MaximizeRespDTO) MaximizeResp {
	return MaximizeResp{
		ReqIDs:      dto.ReqIDs,
		TotalProfit: dto.TotalProfit,
		Max:         dto.Max,
		Min:         dto.Min,
		Avg:         dto.Avg,
	}
}
