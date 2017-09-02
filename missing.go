package panel


func (p Panel) FillMissing() Panel {
	return FillMissing(p)
}

func FillMissing(p Panel) Panel {
	for header, _ := range p {
		p = p.Map(header, func(x interface{}) interface{} {
			switch x.(type) {
			case nil:
				return ""
			default:
				return x
			}
		})
	}

	// for header, series := range p {
	// 	for _, line := range series {
	// 		fmt.Printf("%s: %v: %v\n", header, reflect.TypeOf(line), line)
	// 	}
	// }

	return p
}
