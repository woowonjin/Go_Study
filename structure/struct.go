package main

import "fmt"

type Blog struct {
	id		int
	name	string
	url		string 
}

// func main() {
// 	// case 1
// 	var blog Blog
// 	blog.id = 1
// 	blog.name = "pronist"
// 	blog.url = "httl://pronist.tistory.com"

// 	// case 2
// 	blog = Blog{id: 1, name, "pronist", url: "http://pronist.tistory.com"}

// 	// case 3
// 	blog = Blog{1, "pronist", "http://pronist.tistory.com"}
// }

// func main() {
// 	var blog = struct {
// 		id		int
// 		name	string
// 		url		string
// 	}{1, "pronist", "http://pronist.tistory.com"}

// 	fmt.Println(blog)
// }

func main() {
	fmt.Println(struct{}{})
}