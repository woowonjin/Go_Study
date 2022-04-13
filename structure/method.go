package main

import "fmt"

type Blog struct {
	id		int
	name	string
	url		string
	posts 	[]*Post
}

type Post struct {
	title	string
	content string
}

// b *Blog 가 기존 class의 this 역할
func (b *Blog) write(p *Post) {
	b.posts = append(b.posts, p)
}

func main() {
	blog := Blog{1, "pronist", "http://pronist.tistory.com", []*Post{}}
	blog.write(&Post{"Go: Hello, World", "This is Golang"})
	blog.write(&Post{"Go: White in Go", "Java is verbose, Python is too Slow"})

	for _, post := range blog.posts {
		fmt.Println(*post)
	}
}