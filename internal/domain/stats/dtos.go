package stats

import (
	"time"
)

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
	ReqIDs      []string
	TotalProfit float64
	Max         float64
	Min         float64
	Avg         float64
}

// TODO: refactor to validation and include in the DTO
func toDate(checkIn string, nights int) (time.Time, error) {
	d, err := time.Parse(time.DateOnly, checkIn)
	if err != nil {
		return time.Now(), err
	}
	return d.AddDate(0, 0, nights), nil

}

// areInTime
// TODO: refactor to have less repetitive code and better reading
func areInTime(in []MaximizeReqDTO) (map[string]MaximizeReqDTO, error) {
	intime := make(map[string]MaximizeReqDTO)
	avoid := make(map[string]interface{})
	for _, p := range in {
		if _, ok := avoid[p.ReqID]; ok {
			continue
		}

		pcheckin, err := toDate(p.CheckIn, 0)
		if err != nil {
			return nil, err
		}
		// pcheckout := pcheckin.Add(time.Duration(p.Nights) * (24 * time.Hour))
		pcheckout, err := toDate(p.CheckIn, p.Nights)
		if err != nil {
			return nil, err
		}

		for _, j := range in {
			if j.ReqID == p.ReqID {
				continue
			}

			if _, ok := avoid[j.ReqID]; ok {
				continue
			}

			jcheckin, err := toDate(j.CheckIn, 0)
			if err != nil {
				return nil, err
			}
			jcheckout, err := toDate(j.CheckIn, j.Nights)
			if err != nil {
				return nil, err
			}

			cin := pcheckin.Sub(jcheckin).Abs().Hours()
			cout := pcheckout.Sub(jcheckout).Abs().Hours()

			if cin > 48 && cout > 48 {
				if _, ok := intime[p.ReqID]; !ok {
					intime[p.ReqID] = p
				}
				continue
			}

			if _, ok := avoid[j.ReqID]; !ok {
				avoid[j.ReqID] = struct{}{}
			}
		}

	}
	return intime, nil
}
