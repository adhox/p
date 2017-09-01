package main

func (p Panel) Rows(points ...int) Panel {
	l := len(points)
	switch {
	case l == 1:
		return Rows(p, points[0])
	case l > 1:
		return Rows(p, points[0], points[1])
	default:
		return Rows(p)
	}

}

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

	for header, _ := range p {
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
