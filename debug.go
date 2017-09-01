package main

import (
	"fmt"
	"reflect"
)

// Dtypes displays data types
func (p Panel) Dtypes(heads ...string) {
	Dtypes(p, heads...)
}

// TODO
// limit print to head in heads
func Dtypes(p Panel, heads ...string) {
	for head, col := range p {
		for row, val := range col {
			switch t := val.(type) {
			case float64:
				ref := reflect.TypeOf(t)
				fmt.Printf("head(%s) row(%d) val(%v) ref(%s)\n", head, row, val, ref)
			case int:
				ref := reflect.TypeOf(t)
				fmt.Printf("head(%s) row(%d) val(%v) ref(%s)\n", head, row, val, ref)
			case bool:
				ref := reflect.TypeOf(t)
				fmt.Printf("head(%s) row(%d) val(%v) ref(%s)\n", head, row, val, ref)
			case string:
				ref := reflect.TypeOf(t)
				fmt.Printf("head(%s) row(%d) val(%v) ref(%s)\n", head, row, val, ref)
			default:
				fmt.Printf("head(%s) row(%d) val(%v) ref(%s)\n", head, row, val, reflect.TypeOf(val))
			}
		}
	}
}
