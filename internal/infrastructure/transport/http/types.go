package http

// moove to domain
type statsPayload struct {
	RequestID   string `json:"request_id"`
	CheckIn     string `json:"check_in"`
	Nights      int    `json:"nights"`
	SellingRate int    `json:"selling_rate"`
	Margin      int    `json:"margin"`
}

type Profit struct {
	Avg float64 `json:"avg_night"`
	Min float64 `json:"min_night"`
	Max float64 `json:"max_night"`
}
