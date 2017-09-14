package panel

// Join ...
func Join(p1, p2 Panel, joinType string, on ...string) Panel {
	hash := []map[string]interface{}{}
	p := New(nil)
	headers := []string{}

	// get headers ready
	for _, h := range append(p1.Headers(), p2.Headers()...) {
		if !stringInSlice(h, headers) {
			headers = append(headers, h)
		}
	}

	switch joinType {
	case "inner":
		for r1 := 0; r1 < p1.Size().Length; r1++ {
			for r2 := 0; r2 < p2.Size().Length; r2++ {

				var matches int
				for i := range on {

					if p1[on[i]][r1] == p2[on[i]][r2] {
						matches++
					}
				}

				if len(on) == matches {
					m := map[string]interface{}{}
					for col := range p1 {
						m[col] = p1[col][r1]
					}
					for col := range p2 {
						m[col] = p2[col][r2]
					}
					hash = append(hash, m)
				}
			}
		}

		for _, h := range headers {
			p[h] = []interface{}{}
		}

		for row := 0; row < len(hash); row++ {
			for _, h := range headers {
				p[h] = append(p[h], hash[row][h])
			}
		}
	case "left":
		for r1 := 0; r1 < p1.Size().Length; r1++ {
			joinCount := 0
			for r2 := 0; r2 < p2.Size().Length; r2++ {

				var matches int
				for i := range on {

					if p1[on[i]][r1] == p2[on[i]][r2] {
						matches++
					}
				}
				if len(on) == matches {
					m := map[string]interface{}{}
					for col := range p1 {
						m[col] = p1[col][r1]
					}
					for col := range p2 {
						m[col] = p2[col][r2]
					}
					hash = append(hash, m)
					joinCount++

				}
			}
			if joinCount == 0 {
				m := map[string]interface{}{}
				for col := range p1 {
					m[col] = p1[col][r1]
				}
				for col := range p2 {
					if !stringInSlice(col, on) {
						m[col] = nil
					}
				}
				hash = append(hash, m)
				joinCount++
			}
		}

		for _, h := range headers {
			p[h] = []interface{}{}
		}

		for row := 0; row < len(hash); row++ {
			for _, h := range headers {
				p[h] = append(p[h], hash[row][h])
			}
		}

	case "right":
		return Join(p2, p1, "left", on...)

	case "full":
		m1 := Join(p1, p2, "left", on...)
		m2 := Join(p1, p2, "right", on...)
		return m1.Concat(m2).Unique()

	// // default is rerun as inner join
	default:
		if joinType != "inner" {
			return Join(p1, p2, "inner", on...)
		}
	}
	return p
}

// InnerJoin ...
func (p1 Panel) InnerJoin(p2 Panel, on string) Panel {
	return Join(p1, p2, "inner", on)
}

// LeftJoin ...
func (p1 Panel) LeftJoin(p2 Panel, on string) Panel {
	return Join(p1, p2, "left", on)
}

// RightJoin ...
func (p1 Panel) RightJoin(p2 Panel, on string) Panel {
	return Join(p1, p2, "right", on)
}

// FullJoin ...
func (p1 Panel) FullJoin(p2 Panel, on string) Panel {
	return Join(p1, p2, "full", on)
}
