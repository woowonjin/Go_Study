// Go 에는 while 문이 없다.
package main

import "fmt"

// func main() {
// 	sum := 0
// 	for i := 0; i < 10; i++ {
// 		sum += i
// 	}
// 	fmt.Println(sum)
// }

// func main() {
// 	sum := 1
// 	for sum < 100 {
// 		sum += sum
// 	}
// 	fmt.Println(sum)
// }

// func main() {
// 	for {
// 		fmt.Println("무한 루프 !!")
// 	}
// }

// func main() {
// 	runes := []rune{'A', 'B', 'C'}
// 	for _, r := range runes {
// 		fmt.Println(r)
// 	}
// }

func main() {
	numbers := []int{2, 4, 6, 8, 10, 11, 14}
	// var numbers []int = []int{2, 4, 6, 8, 10, 11, 14}
	for _, num := range numbers {
		if num % 2 == 0 {
			fmt.Println(num)
			continue
		}
		break
	}
}