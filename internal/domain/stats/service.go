package stats

import (
	"fmt"
)

type Service struct{}

func NewService() *Service {
	svc := &Service{}
	return svc
}

// ProfitsPerNight
func (s *Service) ProfitsPerNight(payload []ProfitsPerNightReqDTO) ProfitsPerNightRespDTO {
	var profits []float64
	for _, p := range payload {
		if p.Margin == 0 || p.Nights == 0 {
			profits = append(profits, 0)
			continue
		}

		profit := (p.SellingRate * float64(p.Margin) / 100) / float64(p.Nights)
		profits = append(profits, profit)
	}

	out := newProfit(profits)
	return ProfitsPerNightRespDTO{
		Avg: out.avg,
		Max: out.max,
		Min: out.min,
	}
}

func (s *Service) Maximize(payload []MaximizeReqDTO) (MaximizeRespDTO, error) {
	resp, err := maximize(payload)
	if err != nil {
		return MaximizeRespDTO{}, fmt.Errorf("internal error: %w", err)
	}
	if resp == nil {
		return MaximizeRespDTO{}, fmt.Errorf("internal error: maximize is nil")

	}
	return *resp, nil
}
