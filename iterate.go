package main

func (p Panel) Iter() []map[string]interface{} {
	table := []map[string]interface{}{}

	for i := 0; i < p.Size().Length; i++ {
		row := map[string]interface{}{}
		for col, val := range p {
			row[col] = val[i]
		}
		table = append(table, row)
	}
	return table
}

// func (p Panel) Map2(fn func(int, map[string]interface{}) map[string]interface{}) Panel {
// 	for k, row := range p.Iter() {
// 		fn(k, row)
// 	}
// }
