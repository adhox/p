package panel

// Concat ...
func (p1 Panel) Concat(p2 Panel) Panel {
	return Concat(p1, p2)
}

// Concat ...
func Concat(p1, p2 Panel) Panel {
	for col := range p1 {
		p1[col] = append(p1[col], p2[col]...)
	}
	return p1
}

// simply stack two panels together
// mimics SQL:
//      select 1 as A, 2 as B
//      union all
//      select 3 as A, 4 as B

// TODO
// change Union and UnionAll signatures to something like...
// `func Union(ps ...Panel) Panel`
// where it loops through endless panels
// ... or maybe just implement `Reduce` func
// that could do the same, given a func

// Union ...
func Union(p1, p2 Panel) Panel {
	return p1.Concat(p2).Unique()
}

// UnionAll ...
func UnionAll(p1, p2 Panel) Panel {
	return p1.Concat(p2)
}

// Union ...
func (p1 Panel) Union(p2 Panel) Panel {
	return Union(p1, p2)
}

// UnionAll ...
func (p1 Panel) UnionAll(p2 Panel) Panel {
	return UnionAll(p1, p2)
}
