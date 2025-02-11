package stats

// type ServiceStats interface {
// 	ProfitsPerNiught(xas,xd []float64,xsm
// }

type Service struct {
	// Deppendenciakes here
}

type ProfitsDTO struct {
	Max float64
	Min float64
	Avg float64
}

func (s *Service) ProfitsPerNight(xs []float64) ProfitsDTO {
	p := newProfit(xs)
	return ProfitsDTO{
		Avg: p.avg,
		Max: p.max,
		Min: p.min,
	}
}
