package entity

import "github.com/duhruh/tackle/domain"

type Blog interface {
	Identity() domain.Identity
	Posts() []Post
	Name() string
	SetName(name string)
	SetIdentity(id domain.Identity)
	AddPost(post Post)
}

type blog struct {
	id    domain.Identity
	posts []Post
	name  string
}

func NewBlog() Blog {
	return &blog{}
}

func (b *blog) Identity() domain.Identity {
	return b.id
}

func (b *blog) SetIdentity(id domain.Identity) {
	b.id = id
}

func (b *blog) Posts() []Post {
	return b.posts
}

func (b *blog) AddPost(post Post) {
	b.posts = append(b.posts, post)
}

func (b *blog) Name() string {
	return b.name
}

func (b *blog) SetName(name string) {
	b.name = name
}
