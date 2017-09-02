package panel


func (p Panel) Sum(col string) float64 {
	return Sum(p, col)
}

func Sum(p Panel, col string) (sum float64) {
	switch p[col][0].(type) {
	case int:
		for _, val := range p[col] {
			sum += float64(val.(int))
		}
	case float64:
		for _, val := range p[col] {
			sum += val.(float64)
		}
	}
	return
}
