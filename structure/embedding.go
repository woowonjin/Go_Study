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

type User struct {
	id		int
	email	string
}

type TistoryBlog struct {
	subscribers	[]*User
	Blog
}

func (b* TistoryBlog) Subscribers() []*User {
	return b.subscribers
}

func (u *User) Subscribe(b *TistoryBlog) {
	b.subscribers = append(b.subscribers, u)
}

func (b *Blog) Write(p *Post){
	b.posts = append(b.posts, p)
}

func main() {
	tistoryBlog := TistoryBlog{
		[]*User{},
		Blog{1, "pronist", "http://pronist.tistory.com", []*Post{}},
	}
	helloPost := Post{"Go: Hello, World", "This is Golang"}

	tistoryBlog.Write(&helloPost)

	user := User{1, "wonjin@naver.com"}
	user.Subscribe(&tistoryBlog)

	for _, subscriber := range tistoryBlog.Subscribers() {
		fmt.Println(*subscriber)
	}
}