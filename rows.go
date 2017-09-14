package panel

// Rows ...
func (p Panel) Rows(points ...int) Panel {
	// l := len(points)
	// switch {
	// case l == 1:
	// 	return Rows(p, points[0])
	// case l > 1:
	// 	return Rows(p, points[0], points[1])
	// default:
	// 	return Rows(p)
	// }
	return Rows(p, points...)

}

// Rows ...
func Rows(p Panel, points ...int) Panel {
	var start, end int
	tempPanel := Panel{}
	l := len(points)

	switch {
	case l == 1:
		start = points[0]
		end = -1
	case l > 1:
		start = points[0]
		end = points[1]
	default:
		start = 0
		end = -1
	}

	for header := range p {
		c := []interface{}{}
		if end == -1 {
			for row, val := range p[header] {
				if row >= start {
					// c[row] = val
					c = append(c, val)
				}
			}
			tempPanel.Add(header, c)

		} else {
			for row, val := range p[header] {
				if row >= start && row <= end {
					// c[row] = val
					c = append(c, val)
				}
			}
			tempPanel.Add(header, c)

		}
	}

	return tempPanel
}

// Head ...
func (p Panel) Head(i ...int) Panel {
	return Head(p, i...)
}

// Head ...
func Head(p Panel, i ...int) Panel {
	if len(i) == 0 {
		return p.Rows(0, 4)
	}
	return p.Rows(0, i[0]-1)
}

// Tail ...
func (p Panel) Tail(i ...int) Panel {
	return Tail(p, i...)
}

// Tail ...
func Tail(p Panel, i ...int) Panel {
	ln := p.Size().Length
	if len(i) == 0 {
		return p.Rows(ln-5, ln)
	}
	return p.Rows(ln-i[0], ln)
}
