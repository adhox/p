package main

import "fmt"

func (p Panel) Mode(col string) map[string]int {
	return Mode(p, col)
}

func Mode(p Panel, col string) map[string]int {
	counter := make(map[string]int)
	for _, val := range p[col] {
		vval := fmt.Sprintf("%v", val)
		if counter[vval] > 0 {
			counter[vval] += 1
		} else {
			counter[vval] = 1
		}
	}

	mx := 0
	var top map[string]int

	for i, c := range counter {
		if c > mx {
			mx = c
			top = map[string]int{i: c}
		} else if c == mx {
			top[i] = c
		}
	}
	return top
}
