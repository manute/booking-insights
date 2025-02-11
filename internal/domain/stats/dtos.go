package stats

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
