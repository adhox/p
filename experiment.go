package main

// import "fmt"

// type Panel map[string][]interface{}

// func main() {
// 	people := Panel{}

// 	people.Add("gender", map[int]interface{}{2: "M", 3: "F", 0: "F", 1: "O"})
// 	people.Add("ages", []int{65, 76, 34, 23, 90})
// 	people.Add("names", []string{"Bryan", "Lucy", "Ruben", "Joyce", "Bea"})

// 	fmt.Println(people.Map("ages", func(v int) string {

// 		return fmt.Sprintf("_%d_", v+10)

// 	}))
// }

// func (p Panel) Add(name string, v interface{}) Panel {
// 	switch t := v.(type) {
// 	case []interface{}:
// 		p[name] = t
// 	case []int:
// 		slice := make([]interface{}, len(t))
// 		for k, v := range t {
// 			slice[k] = v
// 		}
// 		p[name] = slice
// 	case []string:
// 		slice := make([]interface{}, len(t))
// 		for k, v := range t {
// 			slice[k] = v
// 		}
// 		p[name] = slice
// 	case []float64:
// 		slice := make([]interface{}, len(t))
// 		for k, v := range t {
// 			slice[k] = v
// 		}
// 		p[name] = slice
// 	}
// 	return p
// }
