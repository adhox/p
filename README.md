# Panel

Data analysis framework for Go.

#### Examples

Use results from a SQL query. 
```go
package main

import (
	"fmt"
	. "github.com/brydavis/panel"
)

func main() {
	query := `select distinct Id, FirstName, LastName from Persons where FirstName like 'A%'`

	conn := Connect("mssql", "config.json") // Connecting to an Azure DB for example 
	persons := Query(conn, query)

	fmt.Printf(p.Select("FirstName").Unique())

	McNames := persons.Map("LastName", func(s string) string {
		return fmt.Sprintf("Mc%s", s)
	})

	fmt.Println(McNames.Select("FirstName", "LastName"))
}
```
Upload a CSV and map a function
```go
	{ // Read a CSV
		df := Read("data/iris.csv", true)
		df = df.Map("species", func(s string) string {
			return fmt.Sprintf("__%s", s)
		})

		fmt.Println(df["species"])

	}
```



#### Features
1. Map
2. Filter
3. Max
4. Min
5. ...



<br>
<br>

<hr>
<small>
<strong>&copy; 2015-2016 [MIT License](https://github.com/openwonk/mit_license)</strong>
</small>
# Info
