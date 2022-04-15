package main

import (
	"fmt"
	"encap/blog"
)

func main() {
	b := blog.NewBlog(1, "pronist", "http://pronist.tistory.com", []*blog.Post{})
	helloPost := blog.NewPost("Go: Hello, World", "This is GoLang")

	helloPost.SetTitle("Go: Hello, Go!")
	b.Write(helloPost)

	for _, post := b.Posts() {
		fmt.Println(*post)
	}
}