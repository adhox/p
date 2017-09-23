package panel

// Panel is a basically tabular-shaped data s
type Panel map[string][]interface{}

// New creates a new Panel
func New(data interface{}) Panel {
	p := make(Panel)
	switch t := data.(type) {
	case CSV:
	case XML:
	case JSON:
	case TSV:
	case SQL:

	case map[string][]interface{}:
		p = Panel(p)
	case map[string][]string:
		for head := range t {
			p.Add(head, t[head])
		}
	case map[string][]bool:
		for head := range t {
			p.Add(head, t[head])
		}

	case map[string][]int:
		for head := range t {
			p.Add(head, t[head])
		}

	case map[string][]float64:
		for head := range t {
			p.Add(head, t[head])
		}

	case [][]string:
		columns, body := t[0], t[1:]
		for row := range body {
			for c, col := range columns {
				p[col] = append(p[col], body[row][c])
			}
		}
	case []map[string]interface{}:
		// ***** TODO: check if interface is string or slice of data???
		// columns := getStringKeys(t)
		// fmt.Println("here")
		// // columns, body := t[0], t[1:]
		// for row := range t {
		// 	for c, col := range columns {
		// 		p[col] = append(p[col], body[row][c])
		// 	}
		// }
	default:

	}

	p = p.Clean()
	Meta = Meta.Append(&p)
	return p

}

// Select ...
func (p Panel) Select(c ...interface{}) Panel {
	cols := []string{}
	columns := p.Columns()

	for i := range c {
		switch t := c[i].(type) {
		case string:
			if t == "*" {
				return p
			}
			cols = append(cols, t)
		case int:
			cols = append(cols, columns[t])
		case []interface{}:
			for x := range t {
				switch tt := t[x].(type) {
				case string:
					if tt == "*" {
						return p
					}
					cols = append(cols, tt)
				case int:
					cols = append(cols, columns[tt])
				}
			}
		case []string:
			for x := range t {
				cols = append(cols, t[x])
			}
		case []int:
			for x := range t {
				cols = append(cols, columns[t[x]])
			}
		}
	}

	if len(cols) == 0 {
		return p
	}

	tempPanel := make(Panel)

	for _, col := range cols {
		if col == "*" {
			p.Clone(tempPanel)
			return tempPanel
		}
		tempPanel.Add(col, p[col])
	}
	return tempPanel
}

// Clone ...
func (p Panel) Clone(dst Panel) {
	for header, series := range p {
		dst.Add(header, series)
	}
}
