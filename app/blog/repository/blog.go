package repository

import (
	"errors"

	"github.com/duhruh/blog/app/blog/entity"
	blogerror "github.com/duhruh/blog/app/blog/errors"

	"github.com/duhruh/blog/app/db"
	"github.com/duhruh/tackle/domain"

	upper "upper.io/db.v3"
)

type BlogRepository interface {
	FindByIdentity(id domain.Identity) (entity.Blog, error)
	Create(b entity.Blog) entity.Blog
	All() ([]entity.Blog, error)
	Update(b entity.Blog) (entity.Blog, error)
}

type blogRepository struct {
	connection db.DatabaseConnection
}

type blog struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

func NewBlogRepository(connection db.DatabaseConnection) BlogRepository {
	return blogRepository{connection: connection}
}

func (br blogRepository) blogTable() upper.Collection {
	return br.connection.Connection().Collection("blog")
}

func (br blogRepository) FindByIdentity(id domain.Identity) (entity.Blog, error) {
	var b blog
	var rb entity.Blog
	rb = entity.NewBlog()

	res := br.blogTable().Find(upper.Cond{"id": id.Identity().(string)})
	err := res.One(&b)

	if err != nil {
		return rb, err
	}

	return br.inflateBlogEntity(rb, b), nil
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
		rb = append(rb, br.inflateBlogEntity(entity.NewBlog(), bb))
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

func (br blogRepository) inflateBlogEntity(b entity.Blog, raw blog) entity.Blog {
	b.SetIdentity(domain.NewIdentity(raw.ID))
	b.SetName(raw.Name)
	return b
}

func (br blogRepository) deflateBlogEntity(b entity.Blog) blog {
	var bb blog
	bb.ID = b.Identity().Identity().(string)
	bb.Name = b.Name()
	return bb
}
