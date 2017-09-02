package panel


import "fmt"

// NOT WORKING

// func (p *Panel) Reset() {
// 	Reset(p)
// }

func Reset(p Panel) {
	tempPanel := make(Panel)
	// for head, col := range *p {
	// 	tempPanel.Add(header, data)
	// }

	p = tempPanel
	// p = &a
	fmt.Printf("\n\n%v\n\n", p)

}
