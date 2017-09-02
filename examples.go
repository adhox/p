package panel


// import (
// 	"fmt"
// )

// // TODO
// // - Transform Go structures into frames

// func main() {
// 	port := 8844

// 	{
// 		df := New(map[string][]string{
// 			"Q": []string{"l", "o", "r", "q"},
// 			"N": []string{"5", "2", "9", "7"},
// 			"M": []string{"false", "true", "true", "false"},
// 		})

// 		df.Dtypes()
// 		fmt.Println(df.Sum("N"))

// 		go df.Serve(port+1, "test1")

// 	}

// 	{
// 		df := Read("./data/gota.csv", true)

// 		subset := df.
// 			Filter("a", func(a string) bool {
// 				return a == "a"
// 			}).
// 			Filter("b", func(b float64) bool {
// 				return b >= 4.0

// 			}).
// 			Filter("d", func(d bool) bool {
// 				return d == true
// 			})

// 			// df.PrintRowTypes()
// 		fmt.Println(subset)

// 	}

// 	{ // Read a CSV file with headers via URL
// 		// link := "https://vincentarelbundock.github.io/Rdatasets/csv/datasets/infert.csv"
// 		link := "https://vincentarelbundock.github.io/Rdatasets/csv/Stat2Data/SATGPA.csv"
// 		// link := "http://www.thestranger.com/seattle/Rss.xml?section=529161"

// 		df := Read(link, true)
// 		fmt.Println(df["gpa"])
// 	}

// 	// { // Read a CSV file with headers stored locally
// 	// 	df := Read("data/iris.csv", true)
// 	// 	df = df.Map("species", func(s string) string {
// 	// 		return fmt.Sprintf("__%s", s)
// 	// 	})
// 	// 	fmt.Println(df["species"][:5])
// 	// 	df.Serve(port, "iris")
// 	// }

// 	{ // Read a XLSX file with headers stored locally
// 		df := Read("data/iris.xlsx", true)
// 		df = df.Map("species", func(s string) string {
// 			return fmt.Sprintf("**%s", s)
// 		})

// 		fmt.Println(df["species"][:5])

// 		df = df.Rename(map[string]string{
// 			"species":     "kind",
// 			"petal.width": "pwidth",
// 			"unnamed: 0":  "",
// 		})

// 		df.Serve(port, "iris")
// 	}

// 	// { // Try reading file that doesn't exist
// 	// 	df := Read("hello.json", false)
// 	// 	if df == nil {
// 	// 		fmt.Println("nothing there")
// 	// 	}
// 	// }

// }
