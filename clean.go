package panel

import (
	"reflect"
	"strconv"
	"time"
)

// Sanitize data types within each series
func (p Panel) Clean(cols ...string) Panel {
	return Clean(p, cols...)
}

func Clean(p Panel, cols ...string) Panel {
	tempPanel := make(Panel) // <- implement goroutine for each column
	if len(cols) == 0 {
		for header, series := range p {
			// tally up most frequent
			// data type per column
			cnt := make(map[string]int)
			for _, val := range series {
				switch t := val.(type) {
				case string:
					if f, err := strconv.ParseFloat(t, 64); err == nil {
						ref := reflect.TypeOf(f).String()
						if _, ok := cnt[ref]; ok {
							cnt[ref] += 1
						} else {
							cnt[ref] = 1
						}
					} else if b, err := strconv.ParseBool(t); err == nil {
						ref := reflect.TypeOf(b).String()
						if _, ok := cnt[ref]; ok {
							cnt[ref] += 1
						} else {
							cnt[ref] = 1
						}
					} else if code, yes := isDate(t); yes {
						format, _ := dateFormat(code)
						d, _ := time.Parse(format, t)
						ref := reflect.TypeOf(d).String()
						if _, ok := cnt[ref]; ok {
							cnt[ref] += 1
						} else {
							cnt[ref] = 1
						}
					} else {
						ref := reflect.TypeOf(val).String()
						if _, ok := cnt[ref]; ok {
							cnt[ref] += 1
						} else {
							cnt[ref] = 1
						}
					}
				default:
					ref := reflect.TypeOf(t).String()
					if _, ok := cnt[ref]; ok {
						cnt[ref] += 1
					} else {
						cnt[ref] = 1
					}
				}
			}

			var mxHead string
			var mxVal int
			for s, i := range cnt {
				if i > mxVal { // || s == "float64" {
					mxHead = s
					mxVal = i
				}
			}

			switch mxHead {
			case "float64":
				for _, val := range series {
					ref := reflect.TypeOf(val)
					var f float64
					if ref.String() == "string" {
						f, _ = strconv.ParseFloat(val.(string), 64)
					} else {
						f = val.(float64)
					}
					tempPanel[header] = append(tempPanel[header], f)
				}
			case "int":
				for _, val := range series {
					i, _ := strconv.Atoi(val.(string))
					tempPanel[header] = append(tempPanel[header], i)
				}
			case "string":
				for _, val := range series {
					tempPanel[header] = append(tempPanel[header], val.(string))
				}
			case "bool":
				for _, val := range series {
					b, _ := strconv.ParseBool(val.(string))
					tempPanel[header] = append(tempPanel[header], b)
				}
			case "time.Time":
				for _, val := range series {
					ref := reflect.TypeOf(val)
					var d time.Time
					if ref.String() == "string" {
						s := val.(string)
						code, _ := isDate(s)
						format, _ := dateFormat(code)
						d, _ = time.Parse(format, s)
					} else {
						d = val.(time.Time)
					}
					tempPanel[header] = append(tempPanel[header], d)
				}
			default:
				for _, val := range series {
					tempPanel[header] = append(tempPanel[header], val)
				}
			}

		}
		return tempPanel

	} else {
		specs := make(map[string]bool, len(cols))
		for _, col := range cols {
			specs[col] = true
		}

		for header, series := range p {
			if specs[header] {
				// tally up most frequent
				// data type per column
				cnt := make(map[string]int)
				for _, val := range series {
					switch t := val.(type) {
					case string:
						if f, err := strconv.ParseFloat(t, 64); err == nil {
							ref := reflect.TypeOf(f).String()
							if _, ok := cnt[ref]; ok {
								cnt[ref] += 1
							} else {
								cnt[ref] = 1
							}
						} else if b, err := strconv.ParseBool(t); err == nil {
							ref := reflect.TypeOf(b).String()
							if _, ok := cnt[ref]; ok {
								cnt[ref] += 1
							} else {
								cnt[ref] = 1
							}
						} else if code, yes := isDate(t); yes {
							format, _ := dateFormat(code)
							d, _ := time.Parse(format, t)
							ref := reflect.TypeOf(d).String()
							if _, ok := cnt[ref]; ok {
								cnt[ref] += 1
							} else {
								cnt[ref] = 1
							}
						} else {
							ref := reflect.TypeOf(val).String()
							if _, ok := cnt[ref]; ok {
								cnt[ref] += 1
							} else {
								cnt[ref] = 1
							}
						}
					default:
						ref := reflect.TypeOf(t).String()
						if _, ok := cnt[ref]; ok {
							cnt[ref] += 1
						} else {
							cnt[ref] = 1
						}
					}
				}

				var mxHead string
				var mxVal int
				for s, i := range cnt {
					if i > mxVal || s == "float64" {
						mxHead = s
						mxVal = i
					}
				}

				switch mxHead {
				case "float64":
					for _, val := range series {
						ref := reflect.TypeOf(val)
						var f float64
						if ref.String() == "string" {
							f, _ = strconv.ParseFloat(val.(string), 64)
						} else {
							f = val.(float64)
						}
						tempPanel[header] = append(tempPanel[header], f)
					}
				case "int":
					for _, val := range series {
						i, _ := strconv.Atoi(val.(string))
						tempPanel[header] = append(tempPanel[header], i)
					}
				case "string":
					for _, val := range series {
						tempPanel[header] = append(tempPanel[header], val.(string))
					}
				case "bool":
					for _, val := range series {
						b, _ := strconv.ParseBool(val.(string))
						tempPanel[header] = append(tempPanel[header], b)
					}
				case "time.Time":
					for _, val := range series {
						ref := reflect.TypeOf(val)
						var d time.Time
						if ref.String() == "string" {
							s := val.(string)
							code, _ := isDate(s)
							format, _ := dateFormat(code)
							d, _ = time.Parse(format, s)
						} else {
							d = val.(time.Time)
						}
						tempPanel[header] = append(tempPanel[header], d)
					}
				default:
					for _, val := range series {
						tempPanel[header] = append(tempPanel[header], val)
					}
				}
			} else {
				tempPanel[header] = series
			}
		}
		return tempPanel
	}
	return p
}
