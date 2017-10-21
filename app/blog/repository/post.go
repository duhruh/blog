package repository

import (
	"github.com/duhruh/blog/app/blog/entity"

	"github.com/duhruh/blog/app/db"
	"github.com/duhruh/tackle/domain"

	"context"
	"github.com/duhruh/blog/app/blog/factory"
	upper "upper.io/db.v3"
)

type PostRepository interface {
	FindByIdentity(id domain.Identity) (entity.Post, error)
	Create(b entity.Post) entity.Post
	All() []entity.Post
	WithContext(ctx context.Context) PostRepository
}

type postRepository struct {
	connection db.DatabaseConnection
	factory    factory.PostFactory
	ctx        context.Context
}

func NewPostRepository(connection db.DatabaseConnection, factory factory.PostFactory) PostRepository {
	return postRepository{connection: connection, ctx: context.Background(), factory: factory}
}

func (br postRepository) postTable() upper.Collection {
	return br.connection.ConnectionWithContext(br.ctx).Collection("post")
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

	return br.factory.PostFromImmutable(b), nil
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
		rb = append(rb, br.factory.PostFromImmutable(bb))
	}

	return rb
}

func (br postRepository) deflatePostEntity(b entity.Post) post {
	var bb post
	bb.ID = b.Identity().Identity().(string)
	bb.B = b.Body()
	bb.Bid = b.BlogId().Identity().(string)
	return bb
}

func (br postRepository) WithContext(ctx context.Context) PostRepository {
	br.ctx = ctx
	return br
}
