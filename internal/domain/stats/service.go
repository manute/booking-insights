package stats

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
