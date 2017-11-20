package repository

import (
	"github.com/duhruh/tackle/domain"

	"github.com/duhruh/blog/app/blog/entity"
)

type blog struct {
	ID string `db:"id"`
	Na string `db:"name"`
}

func (b blog) Identity() domain.Identity {
	return domain.NewIdentity(b.ID)
}
func (b blog) Name() string {
	return b.Na
}
func (b blog) Posts() []entity.Post {
	var p []entity.Post
	return p
}

type post struct {
	ID  string `db:"id"`
	B   string `db:"body"`
	Bid string `db:"blog_id"`
}

func (p post) Identity() domain.Identity { return domain.NewIdentity(p.ID) }
func (p post) Body() string              { return p.B }
func (p post) BlogId() domain.Identity   { return domain.NewIdentity(p.Bid) }
