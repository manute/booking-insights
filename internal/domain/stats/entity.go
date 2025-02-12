package stats

// Profits per night
type profits struct {
	avg float64
	min float64
	max float64
}

func newProfit(xs []float64) profits {
	return profits{
		avg: avg(xs),
		min: min(xs),
		max: max(xs),
	}
}

// Returns average of a series of numbers
func avg(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	if total == 0 {
		return 0
	}
	return total / float64(len(xs))
}

// Returns smallest number in a series
func min(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}

	min := xs[0]
	for _, x := range xs {
		if min > x {
			min = x
		}
	}
	return min
}

// Returns largest number in a series
func max(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}

	max := xs[0]
	for _, x := range xs {
		if max < x {
			max = x
		}
	}
	return max
}

func maximize(req []MaximizeReqDTO) (*MaximizeRespDTO, error) {
	in, err := areInTime(req)
	if err != nil {
		return nil, err
	}

	res := &MaximizeRespDTO{ReqsID: []string{}, TotalProfit: 0}
	var profits []float64
	for id, p := range in {
		if p.Margin == 0 || p.Nights == 0 {
			profits = append(profits, 0)
			continue
		}

		b := p.SellingRate * float64(p.Margin) / 100
		profit := b / float64(p.Nights)

		profits = append(profits, profit)
		res.ReqsID = append(res.ReqsID, id)
		res.TotalProfit += b
	}

	pfpn := newProfit(profits)
	res.Avg = pfpn.avg
	res.Max = pfpn.max
	res.Min = pfpn.min

	return res, nil
}
