package stats

type ProfitsPerNightReqDTO struct {
	SellingRate float64
	Margin      int
	Nights      int
}

type ProfitsPerNightRespDTO struct {
	Max float64
	Min float64
	Avg float64
}

// TODO: Merge with the above ?
type MaximizeReqDTO struct {
	ReqID       string
	CheckIn     string
	SellingRate float64
	Margin      int
	Nights      int
}

type MaximizeRespDTO struct {
	ReqsID      []string
	TotalProfit float64
	Max         float64
	Min         float64
	Avg         float64
}
