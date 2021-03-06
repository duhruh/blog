package entity

import (
	"github.com/duhruh/tackle/domain"
)

type ImmutablePost interface {
	Identity() domain.Identity
	Body() string
	BlogId() domain.Identity
}
type Post interface {
	ImmutablePost
	SetIdentity(id domain.Identity)
	SetBody(body string)
	SetBlog(blog Blog)
}
type post struct {
	id     domain.Identity
	body   string
	blogId domain.Identity
}

func NewPost() Post {
	return &post{}
}

func (p *post) Identity() domain.Identity {
	return p.id
}

func (p *post) SetIdentity(id domain.Identity) {
	p.id = id
}

func (p *post) Body() string {
	return p.body
}

func (p *post) SetBody(body string) {
	p.body = body
}

func (p *post) SetBlog(blog Blog) {
	p.blogId = blog.Identity()
}

func (p *post) BlogId() domain.Identity {
	return p.blogId
}
