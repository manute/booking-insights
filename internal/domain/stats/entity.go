package stats

// Prodfits per nnight
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
