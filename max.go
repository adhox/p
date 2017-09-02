package panel

func (p Panel) Max(col string) interface{} {
	return Max(p, col)
}

func Max(p Panel, col string) interface{} {
	// fmt.Printf("%v => %T\n", p[col][0], p[col][0])
	switch a := p[col][0].(type) {
	case int:
		var mx float64
		for _, val := range p[col] {
			if b := float64(val.(int)); b > mx {
				mx = b
			}
		}
		return mx
	case float64:
		var mx float64
		for _, val := range p[col] {
			if b := val.(float64); b > mx {
				mx = b
			}
		}
		return mx

	case string:
		mx := a
		ln := len(a)
		for _, val := range p[col][1:] {
			if b := len(val.(string)); b > ln {
				mx = val.(string)
				ln = b
			}
		}
		return mx
	}

	return nil
}
