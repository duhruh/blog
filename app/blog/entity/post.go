package entity

import "github.com/duhruh/tackle/domain"

type Post interface {
	Identity() domain.Identity
	Body() string
	SetIdentity(id domain.Identity)
	SetBody(body string)
}
type post struct {
	id   domain.Identity
	body string
}

func NewPost() Post {
	return post{}
}

func (p post) Identity() domain.Identity {
	return p.id
}

func (p post) SetIdentity(id domain.Identity) {
	p.id = id
}

func (p post) Body() string {
	return p.body
}

func (p post) SetBody(body string) {
	p.body = body
}
