package main

func (p Panel) Range(col string) []interface{} {
	return Range(p, col)
}

func Range(p Panel, col string) []interface{} {
	return []interface{}{p.Min(col), p.Max(col)}
}
