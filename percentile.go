package main

// import (
// 	"fmt"
// 	"sort"
// )

// // Percentiles
// func (p Panel) Percentile(col string) Panel {
// 	// TODO
// 	// - JOIN percentiles with associated values in Panel
// 	p.AddSeries(fmt.Sprintf("%s_percentiles", col), Percentile(p[col]))
// 	return p
// }

// func Rank(nums interface{}) map[interface{}]int {
// 	switch t := nums.(type) {
// 	case []int:
// 		m := make(map[interface{}]int, len(t))
// 		sort.Ints(t)
// 		for rank, num := range t {
// 			m[num] = rank + 1
// 		}
// 		return m
// 	case []string:
// 		m := make(map[interface{}]int, len(t))
// 		sort.Strings(t)
// 		for rank, num := range t {
// 			m[num] = rank + 1
// 		}
// 		return m
// 	default:
// 		d := make(map[interface{}]int)
// 		return d
// 	}
// }

// func Percentile(nums interface{}) map[interface{}]float64 {

// 	switch t := nums.(type) {
// 	case []string:
// 		sort.Strings(t)
// 		unitSize := 1.0 / float64(len(t))
// 		freqs := Frequency(t)
// 		m := make(map[interface{}]float64, len(t))

// 		for below, val := range t {
// 			sizeBelow := unitSize * float64(below)
// 			midSizeVal := float64(freqs[val]) * unitSize * 0.5
// 			percentRank := sizeBelow + midSizeVal
// 			if _, exists := m[val]; m[val] > percentRank || !exists {

// 				m[val] = percentRank
// 			}
// 		}
// 		return m

// 	case []int:
// 		sort.Ints(t)
// 		unitSize := 1.0 / float64(len(t))
// 		freqs := Frequency(t)
// 		m := make(map[interface{}]float64, len(t))

// 		for below, val := range t {
// 			sizeBelow := unitSize * float64(below)
// 			midSizeVal := float64(freqs[val]) * unitSize * 0.5
// 			percentRank := sizeBelow + midSizeVal
// 			if _, exists := m[val]; m[val] > percentRank || !exists {

// 				m[val] = percentRank
// 			}
// 		}
// 		return m
// 	}
// 	return make(map[interface{}]float64)
// }

// func Frequency(nums interface{}) map[interface{}]int {
// 	switch t := nums.(type) {
// 	case []int:
// 		cnt := make(map[interface{}]int, len(t))
// 		for _, v := range t {
// 			if val, exists := cnt[v]; exists {
// 				cnt[v] = val + 1
// 			} else {
// 				cnt[v] = 1
// 			}
// 		}
// 		return cnt

// 	}

// 	return make(map[interface{}]int)

// }
