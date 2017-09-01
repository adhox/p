package main

// simply stack two panels together
// mimics SQL:
//      select 1 as A, 2 as B
//      union all
//      select 3 as A, 4 as B

// func Union(ps ...Panel) Panel
// func UnionAll(ps ...Panel) Panel

// func (p Panel) Union(ps ...Panel) Panel {
//     ps = append(ps, p)
//     return Union(ps)
// }

// func (p Panel) UnionAll(ps ...Panel) Panel {
//     ps = append(ps, p)
//     return UnionAll(ps)
// }
