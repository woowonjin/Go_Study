package blog

type Blog struct {
	id		int
	name	string
	url		string
	posts	[]*Post
}

type Post struct {
	title	string
	content	string
}

func NewBlog(id int, name string, url string, posts []*Post) *Blog {
	return &Blog{id, name, url, posts}
}

func NewPost(title string, content string) *Post {
	return &Post{title, content}
}

// Setter
func (p *Post) SetTitle(title string) {
	p.title = title
}

// Getter
func (b *Blog) Posts() []*Post {
	return b.posts
}

func (b* Blog) Write(p *Post) {
	b.posts = append(b.posts, p)
}