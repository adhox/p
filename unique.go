package panel

import (
	"fmt"
	"strings"
)

// func main() {
// 	people := Panel{}

// 	people.Add("gender", map[int]interface{}{2: "M", 3: "F", 0: "M", 1: "O", 4: "F", 5: "M", 6: "F"})
// 	people.Add("ages", []int{65, 76, 34, 23, 90, 65, 55})
// 	people.Add("iqs", []float64{101.15, 125.09, 95.67, 105.81, 99.19, 70.9, 180.0})
// 	people.Add("names", []string{"Bryan", "Lucy", "Ruben", "Joyce", "Bea", "Bryan", "Bea"})

// 	fmt.Println(people)

// 	p := people.Select("names", "gender", "ages").Unique()
// 	fmt.Println(p)
// }

func (p Panel) Distinct() Panel {
	return Unique(p)
}

func (p Panel) Unique() Panel {
	return Unique(p)
}

func Unique(p Panel) Panel {
	var tempVal []string
	var tempCol []string

	for col, srs := range p {
		tempCol = append(tempCol, col)

		for row, val := range srs {
			lenTempVal := len(tempVal)

			if row >= lenTempVal {
				tempVal = append(tempVal, fmt.Sprintf("%v", val))
			} else {
				tempVal[row] += fmt.Sprintf("|%v", val)

			}
		}
	}

	m := make(map[string]bool)

	for _, row := range tempVal {
		m[row] = true
	}

	var final [][]string

	for val, _ := range m {
		final = append(final, strings.Split(val, "|"))
	}

	pnl := Panel{}

	for colNum, colName := range tempCol {
		for _, rowVals := range final {

			pnl[colName] = append(pnl[colName], rowVals[colNum])
		}
	}

	return pnl
}

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

// func (p Panel) Clone(dst Panel) {
// 	for header, series := range p {
// 		dst.Add(header, series)
// 	}
// }
