package main

import "fmt"

// func main() {
// 	// numbers := [5]int{0, 1, 2, 3, 4}
// 	// runes := [5]rune{
// 	// 	'A',
// 	// 	'B',
// 	// 	'C',
// 	// 	'D', // -> 여러줄로 적으려면 마지막에 반드시 콤마를 찍어줄 필요가있음
// 	// }
	
// 	var numbers [5]int // 초기화 안해주면 int는 자동으로 0으로 초기화
// 	for _, num := range numbers {
// 		fmt.Println(num)
// 	}

// }

// func main() {
// 	// numbers := make([]int, 5)
// 	// numbers := []int{0, 1, 2, 3, 4} // 동적 배열
// 	var numbers []int
// 	fmt.Println(len(numbers))
// }

func main() {
	// Slicing
	// numbers := [5]int{0, 1, 2, 3, 4}
	// for _, num := range numbers[1:4] {
	// 	fmt.Println(num)
	// }	
	// s := numbers[1:4]
	// s[0] = 100
	// s[1] = 200
	// fmt.Println(s, numbers)

	// numbers := make([]int, 0)
	// numbers := []int{}
	// for _, num := range [5]int{0, 1, 2, 3, 4} {
	// 	numbers = append(numbers, num)
	// }
	// fmt.Println(numbers)
	
	// s := numbers[1:4]
	// fmt.Println(s)
	// s = append(s, 500)

	// fmt.Println(s, numbers)
	numbers := []int{0, 1, 2, 3, 4}
	s := append(numbers, 500)
	fmt.Println(s, numbers)


}


