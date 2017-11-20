package repository

import (
	"context"
	"errors"

	"github.com/duhruh/tackle/domain"
	upper "upper.io/db.v3"

	"github.com/duhruh/blog/app/blog/entity"
	"github.com/duhruh/blog/app/blog/factory"
	"github.com/duhruh/blog/app/db"
	blogerror "github.com/duhruh/blog/app/errors"
)

type BlogRepository interface {
	FindByIdentity(id domain.Identity) (entity.Blog, error)
	Create(b entity.Blog) entity.Blog
	All() ([]entity.Blog, error)
	Update(b entity.Blog) (entity.Blog, error)
	WithContext(ctx context.Context) BlogRepository
}

type blogRepository struct {
	connection db.DatabaseConnection
	factory    factory.BlogFactory
	ctx        context.Context
}

func NewBlogRepository(connection db.DatabaseConnection, factory factory.BlogFactory) BlogRepository {
	return blogRepository{connection: connection, ctx: context.Background(), factory: factory}
}

func (br blogRepository) blogTable() upper.Collection {
	return br.connection.ConnectionWithContext(br.ctx).Collection("blog")
}

func (br blogRepository) FindByIdentity(id domain.Identity) (entity.Blog, error) {
	var b blog

	res := br.blogTable().Find(upper.Cond{"id": id.Identity().(string)})
	err := res.One(&b)
	if err != nil {
		var rb entity.Blog
		return rb, err
	}

	return br.factory.BlogFromImmutable(b), nil
}

func (br blogRepository) Create(b entity.Blog) entity.Blog {
	bb := br.deflateBlogEntity(b)

	_, err := br.blogTable().Insert(bb)

	if err != nil {
		panic(err) //nooooo
	}

	return b
}

func (br blogRepository) All() ([]entity.Blog, error) {
	var b []blog
	var rb []entity.Blog

	res := br.blogTable().Find()
	err := res.All(&b)

	if err != nil {
		return rb, blogerror.New(blogerror.ErrorCouldNotRetrieveBlogs)
	}

	for _, bb := range b {
		rb = append(rb, br.factory.BlogFromImmutable(bb))
	}

	return rb, nil
}

func (br blogRepository) Update(b entity.Blog) (entity.Blog, error) {
	var bb blog

	res := br.blogTable().Find(upper.Cond{"id": b.Identity().Identity().(string)})
	err := res.One(&bb)

	bb = br.deflateBlogEntity(b)

	err = res.Update(bb)

	if err != nil {
		return b, err
	}

	var nope entity.Blog

	return nope, errors.New("blog not found")
}

func (br blogRepository) deflateBlogEntity(b entity.Blog) blog {
	var bb blog
	bb.ID = b.Identity().Identity().(string)
	bb.Na = b.Name()
	return bb
}

func (br blogRepository) WithContext(ctx context.Context) BlogRepository {
	br.ctx = ctx
	return br
}
