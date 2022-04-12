package main

import (
	"fmt"
	"net/http"
	// "time"
)

// func main() {
// 	response, _ := http.Get("https://example.com")
// 	switch response.StatusCode {
// 	case 200:
// 		fmt.Println("Successful")
// 		fallthrough
// 	case 404:
// 		fmt.Println("Not Found")
// 		fallthrough
// 	default:
// 		fmt.Println("Oops, What is this?")
// 	}
// }

// func main() {
// 	t := time.Now()
// 	switch {
// 	case t.Hour() < 12:
// 		fmt.Println("Good morning!!")
// 	case t.Hour() < 17:
// 		fmt.Println("Good afternoon")
// 	default:
// 		fmt.Println("Good evening")
// 	}
// }

func main() {
	switch response, _ := http.Get("https://example.com"); response.StatusCode {
	case 200:
		fmt.Println("Successful")
	}
}