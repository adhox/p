package main

func (p Panel) Count(col string) float64 {
	return Count(p, col)
}

func Count(p Panel, col string) float64 {
	return float64(len(p[col]))
}
