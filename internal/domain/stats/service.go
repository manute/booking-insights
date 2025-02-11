package stats

type Service struct {
}

type ProfitsPerNightRespDTO struct {
	Max float64
	Min float64
	Avg float64
}

type ProfitsPerNightReqDTO struct {
	SellingRate float64
	Margin      int
	Nights      int
}

// func (s *Service) ProfitsPerNight(xs []float64) ProfitsPerNightRespDTO {
func (s *Service) ProfitsPerNight(payload []ProfitsPerNightReqDTO) ProfitsPerNightRespDTO {
	var data []float64

	for _, p := range payload {
		if p.Margin == 0 || p.Nights == 0 {
			data = append(data, 0)
			continue
		}
		p := (p.SellingRate * float64(p.Margin) / 100) / float64(p.Nights)
		data = append(data, p)
	}

	out := newProfit(data)
	return ProfitsPerNightRespDTO{
		Avg: out.avg,
		Max: out.max,
		Min: out.min,
	}
}
