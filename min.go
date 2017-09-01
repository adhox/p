package main

func (p Panel) Min(col string) interface{} {
	return Min(p, col)
}

func Min(p Panel, col string) interface{} {
	switch a := p[col][0].(type) {
	case int:
		mn := float64(a)
		for _, val := range p[col][1:] {
			if b := float64(val.(int)); b < mn {
				mn = b
			}
		}
		return mn
	case float64:
		mn := a
		for _, val := range p[col][1:] {
			if b := val.(float64); b < mn {
				mn = b
			}
		}
		return mn
	case string:
		mn := a
		ln := len(a)
		for _, val := range p[col][1:] {
			if b := len(val.(string)); b < ln {
				mn = val.(string)
				ln = b
			}
		}
		return mn
	default:
		// return

	}

	return nil
}
