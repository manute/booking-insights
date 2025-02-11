package http

// moove to domain
type statsPayload struct {
	RequestID   string `json:"request_id"`
	CheckIn     string `json:"check_in"`
	Nights      int    `json:"nights"`
	SellingRate int    `json:"sellling_rate"`
	Margin      int    `json:"margin"`
}

type statsPayloadWrapper struct {
	stats []statsPayload
}

type Profit struct {
	Avg float64 `json:"avg_night"`
	Min float64 `json:"min_night"`
	Max float64 `json:"max_night"`
}

func PerNight(payload []statsPayload) Profit {

	// (selling_rate * margin (percentage)) / nights
	var profits []float64
	for _, p := range payload {
		profit := (p.SellingRate * p.Margin / 100) / p.Nights
		profits = append(profits, float64(profit))
	}

	return Profit{
		// Avg: stats.Average(profits),
		// Min: stats.Min(profits),
		// Max: stats.Max(profits),
	}
}
