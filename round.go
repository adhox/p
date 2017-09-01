package main

import (
	"math"
	"strconv"
)

func (p Panel) Round(col string, places ...int) Panel {
	l := len(places)
	switch {
	case l > 0:
		for row, val := range p[col] {
			p[col][row] = Round(val, places[0])
		}
	default:
		for row, val := range p[col] {
			p[col][row] = Round(val, 0)
		}
	}
	return p
}

func Round(v interface{}, places ...int) float64 {
	pl := 0
	if len(places) > 0 {
		pl = places[0]
	}

	shift := math.Pow(10, float64(pl))

	switch t := v.(type) {
	case int:
		return float64(t)
	case string:
		pf, _ := strconv.ParseFloat(t, 64)
		f := math.Floor((pf * shift) + 0.5)
		return f / shift
	default:
		f := math.Floor((v.(float64) * shift) + 0.5)
		return f / shift
	}

}
