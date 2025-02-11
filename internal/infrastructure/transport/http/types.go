package http

type statsPayload struct {
	RequestID   string `json:"request_id"`
	CheckIn     string `json:"check_in"`
	Nights      int    `json:"nights"`
	SellingRate int    `json:"selling_rate"`
	Margin      int    `json:"margin"`
}

type statsResp struct {
	Max float64 `json:"max_night"`
	Min float64 `json:"min_night"`
	Avg float64 `json:"avg_night"`
}
