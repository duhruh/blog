package repository

import (
	"github.com/duhruh/blog/app/blog/entity"

	"github.com/duhruh/blog/app/db"
	"github.com/duhruh/tackle/domain"

	upper "upper.io/db.v3"
)

type PostRepository interface {
	FindByIdentity(id domain.Identity) (entity.Post, error)
	Create(b entity.Post) entity.Post
	All() []entity.Post
}

type postRepository struct {
	connection db.DatabaseConnection
}

type post struct {
	ID     string `db:"id"`
	Body   string `db:"body"`
	BlogId string `db:"blog_id"`
}

func NewPostRepository(connection db.DatabaseConnection) PostRepository {
	return postRepository{connection: connection}
}

func (br postRepository) postTable() upper.Collection {
	return br.connection.Connection().Collection("post")
}

func (br postRepository) FindByIdentity(id domain.Identity) (entity.Post, error) {
	var b post
	var rb entity.Post
	rb = entity.NewPost()

	res := br.postTable().Find(upper.Cond{"id": id.Identity().(string)})
	err := res.One(&b)

	if err != nil {
		return rb, err
	}

	return br.inflatePostEntity(rb, b), nil
}

func (br postRepository) Create(b entity.Post) entity.Post {
	bb := br.deflatePostEntity(b)

	_, err := br.postTable().Insert(bb)

	if err != nil {
		panic(err) //nooooo
	}

	return b
}

func (br postRepository) All() []entity.Post {
	var b []post
	var rb []entity.Post

	res := br.postTable().Find()
	err := res.All(&b)

	if err != nil {
		return rb
	}

	for _, bb := range b {
		rb = append(rb, br.inflatePostEntity(entity.NewPost(), bb))
	}

	return rb
}

func (br postRepository) inflatePostEntity(b entity.Post, raw post) entity.Post {
	b.SetIdentity(domain.NewIdentity(raw.ID))
	b.SetBody(raw.Body)
	p := entity.NewBlog()
	p.SetIdentity(domain.NewIdentity(raw.BlogId))
	b.SetBlog(p)
	return b
}

func (br postRepository) deflatePostEntity(b entity.Post) post {
	var bb post
	bb.ID = b.Identity().Identity().(string)
	bb.Body = b.Body()
	bb.BlogId = b.BlogId().Identity().(string)
	return bb
}
