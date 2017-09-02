package panel


type Size struct {
	Length,
	Width int
}

func (p Panel) Size() Size {
	size := Size{}
	for _, col := range p {
		size.Width += 1
		if len(col) > size.Length {
			size.Length = len(col)
		}
	}
	return size
}
