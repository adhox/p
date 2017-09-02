package panel


// Panel is a basically tabular-shaped data s
type Panel map[string][]interface{}

// New creates a new Panel
func New(data interface{}) Panel {
	p := make(Panel)
	switch t := data.(type) {
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
	default:

	}
	return p.Clean()

}

// Select ...
func (p Panel) Select(cols ...string) Panel {
	tempPanel := make(Panel)
	if len(cols) == 0 || cols[0] == "*" {
		p.Clone(tempPanel)
	} else {
		for _, c := range cols {
			tempPanel.Add(c, p[c])
		}
	}
	return tempPanel
}

// Clone ...
func (p Panel) Clone(dst Panel) {
	for header, series := range p {
		dst.Add(header, series)
	}
}
