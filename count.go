package panel

// Count returns the length of a column
func (p Panel) Count(col string) float64 {
	return Count(p, col)
}

// Count returns the length of a column
func Count(p Panel, col string) float64 {
	return float64(len(p[col]))
}
