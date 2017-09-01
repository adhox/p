package main

// May modify with delimit options: func (p Panel) String(delim string) string
// For example, p.String(",")... How would you pass delim in fmt.Println(p)???
//

import "fmt"
import "sort"

// Print .... (placeholder for pretty printing; undone)
func (p Panel) Print() {
	// Pretty way????
	fmt.Println(p)
}

// String returns printable text
func (p Panel) String() string {
	length := p.Size().Length

	var (
		res  string
		cols []string
	)

	// collect headers
	for col := range p {
		cols = append(cols, col)
	}

	// sort headers
	sort.Strings(cols)

	// add headers to output
	for _, col := range cols {
		res += fmt.Sprintf("%v\t", col)
	}
	res += "\n"

	// for each row
	for row := 0; row < length; row++ {
		s := ""
		for _, col := range cols {
			s += fmt.Sprintf("%v\t", p[col][row])
		}
		res += fmt.Sprintf("%s\n", s)
	}

	return res
}

// HTML returns markup table
func (p Panel) HTML() string {
	length := p.Size().Length

	var (
		res  string
		cols []string
	)

	// collect headers
	for col := range p {
		cols = append(cols, col)
	}

	// sort headers
	sort.Strings(cols)

	// add headers to output
	res = "<table><thead><tr>"
	for _, col := range cols {
		res += fmt.Sprintf("<th>%v</th>", col)
	}
	res += "</tr></thead><tbody>"

	// for each row
	for row := 0; row < length; row++ {
		s := "<tr>"
		for _, col := range cols {
			if col != "" {
				s += fmt.Sprintf("<td contenteditable=\"true\">%v</td>", p[col][row])
			} else {
				s += fmt.Sprintf("<td><strong>%v</strong></td>", p[col][row])
			}
		}
		res += fmt.Sprintf("%s</tr>", s)
	}

	return res + `</tbody></table>
	<style>
		table {
			font-family: arial, sans-serif;
			border-collapse: collapse;
			width: 100%;
		}
		
		td, th {
			border: 1px solid #dddddd;
			text-align: center;
			padding: 8px;
		}
		
		tr:nth-child(even) {
			background-color: #dddddd;
		}
	</style>
	<script>
		var updateButton = document.getElementById("update")
	</script>
	
	
	`
}
