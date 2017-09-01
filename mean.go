package main

func (p Panel) Mean(col string) float64 {
	return Mean(p, col)
}

func Mean(p Panel, col string) float64 {
	return p.Sum(col) / p.Count(col)
}
