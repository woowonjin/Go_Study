package main

import "fmt"

// func main() {
// 	// var m map[string]string = make(map[string]string) // 동적 map
// 	// m := map[string]string{"sayHello": "Hello, Go!"} // literal map

// 	var m map[string]string

// 	fmt.Println(m == nil)

// 	m = map[string]string{"sayHello": "Hello, Go!"}

// 	fmt.Printf("%#v", m["Java"])

// }

// func main() {
// 	m := map[string]string{"sayHello": "Hello, Go!"}
// 	m["java"] = "verbose"

// 	fmt.Println(m)
// 	delete(m, "Java")
// }

// func main() {
// 	m := make(map[string]string)
// 	if _, ok := m["Java"]; ok {
// 		fmt.Printf("%#v", m["Java"])
// 	}
// }

func main() {
	m := map[int]rune{0: 'A', 1: 'B', 2: 'C'}
	for _, r := range m {
		fmt.Println(r)
	}
}