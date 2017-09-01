package main

// Columns ...
func (p Panel) Columns(cols map[string]string) Panel {
	return Columns(p, cols)
}

// Columns ...
func Columns(p Panel, cols map[string]string) Panel {
	if cols != nil {
		tempPanel := make(Panel)
		for header, series := range p {
			if val, ok := cols[header]; ok {
				tempPanel.Add(val, series)
			} else {
				tempPanel.Add(header, series)
			}
		}
		return tempPanel
	}
	return p
}

// Add a column with data
func Add(p Panel, header string, data interface{}) Panel {
	return p.Add(header, data)
}

// Add a column with data
func (p Panel) Add(header string, data interface{}) Panel {
	switch t := data.(type) {
	case []interface{}:
		p[header] = t
	case []int:
		slice := make([]interface{}, len(t))
		for k, v := range t {
			slice[k] = v
		}
		p[header] = slice
	case []string:
		slice := make([]interface{}, len(t))
		for k, v := range t {
			slice[k] = v
		}
		p[header] = slice
	case []float64:
		slice := make([]interface{}, len(t))
		for k, v := range t {
			slice[k] = v
		}
		p[header] = slice
	case map[int]interface{}:
		slice := make([]interface{}, len(t))
		for k, v := range t {
			slice[k] = v
		}
		p[header] = slice

	}
	// todo: add map[int]interface...
	return p
}

// Rename a column
func (p Panel) Rename(cols map[string]string) Panel {
	return Rename(p, cols)
}

// Rename a column
func Rename(p Panel, cols map[string]string) Panel {
	removals := []string{}
	if cols != nil {
		for from, to := range cols {
			p.Add(to, p[from])
			removals = append(removals, from)
		}
	}
	return p.Remove(removals...)
}

// Remove an entire column
func (p Panel) Remove(cols ...string) Panel {
	return Remove(p, cols...)
}

// Remove an entire column
func Remove(p Panel, cols ...string) Panel {
	if len(cols) != 0 {
		for _, col := range cols {
			delete(p, col)
		}
	}
	return p
}
