package main

import "fmt"

// func main() {

// 	p1 := Panel{}

// 	p1.Add("gender", map[int]interface{}{2: "M", 3: "F", 0: "F", 1: "O"})
// 	p1.Add("ages", []int{65, 76, 34, 23, 90})
// 	p1.Add("names", []string{"Bryan", "Lucy", "Ruben", "Joyce", "Bea"})

// 	p2 := Panel{}

// 	p2.Add("iqs", []float64{101.15, 125.09, 95.67, 105.81, 99.19})
// 	p2.Add("names", []string{"Bryan", "Lucy", "Ruben", "Joyce", "Bea"})

// 	p3 := Join(p1, p2, "inner", "on1", "on2")

// 	fmt.Println(p3)

// }

func Join(p1, p2 Panel, join, on1, on2 string) Panel {
	temp1 := p1.Select()

	for col2, srs2 := range p2 {

		temp1.Add(fmt.Sprintf("%s_2", col2), srs2)

	}
	return temp1
}

// func (p Panel) Select(cols ...string) Panel {
// 	tempPanel := make(Panel)
// 	if len(cols) == 0 || cols[0] == "*" {
// 		p.Clone(tempPanel)
// 	} else {
// 		for _, c := range cols {
// 			tempPanel.Add(c, p[c])
// 		}
// 	}
// 	return tempPanel
// }

// type Panel map[string][]interface{}

// func (p Panel) Add(header string, data interface{}) Panel {
// 	switch t := data.(type) {
// 	case []interface{}:
// 		p[header] = t
// 	case []int:
// 		slice := make([]interface{}, len(t))
// 		for k, v := range t {
// 			slice[k] = v
// 		}
// 		p[header] = slice
// 	case []string:
// 		slice := make([]interface{}, len(t))
// 		for k, v := range t {
// 			slice[k] = v
// 		}
// 		p[header] = slice
// 	case []float64:
// 		slice := make([]interface{}, len(t))
// 		for k, v := range t {
// 			slice[k] = v
// 		}
// 		p[header] = slice
// 	case map[int]interface{}:
// 		slice := make([]interface{}, len(t))
// 		for k, v := range t {
// 			slice[k] = v
// 		}
// 		p[header] = slice

// 	}
// 	// todo: add map[int]interface...
// 	return p
// }

func (p1 Panel) InnerJoin(p2 Panel, on1, on2 string) Panel {
	return Join(p1, p2, "inner", on1, on2)

}

func (p1 Panel) LeftJoin(p2 Panel, on1, on2 string) Panel {
	return Join(p1, p2, "left", on1, on2)

}

func (p1 Panel) RightJoin(p2 Panel, on1, on2 string) Panel {
	return Join(p1, p2, "right", on1, on2)

}

func (p1 Panel) FullJoin(p2 Panel, on1, on2 string) Panel {
	return Join(p1, p2, "full", on1, on2)

}
