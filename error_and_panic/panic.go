package main

import "fmt"

// func main() {
// 	var numbers [1]int
// 	for i := 0; i <= len(numbers); i++ {
// 		fmt.Println(numbers[i])
// 	}
// }

func main() {
	defer fmt.Println("Hello, Go!")

	panic("PANIC")
	fmt.Println("Here")
}