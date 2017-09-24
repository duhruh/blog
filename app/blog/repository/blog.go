package repository

import (
	"errors"

	"github.com/duhruh/blog/app/blog/entity"

	"github.com/duhruh/tackle/domain"
)

type BlogRepository interface {
	FindByIdentity(id domain.Identity) (entity.Blog, error)
	Create(b entity.Blog) entity.Blog
	All() []entity.Blog
	Update(b entity.Blog) (entity.Blog, error)
}

type blogRepository struct {
}

var blogs map[string]entity.Blog

func NewBlogRepository() BlogRepository {
	blogs = make(map[string]entity.Blog)
	return blogRepository{}
}

func (br blogRepository) FindByIdentity(id domain.Identity) (entity.Blog, error) {
	return blogs[id.Identity().(string)], nil
}

func (br blogRepository) Create(b entity.Blog) entity.Blog {
	blogs[b.Identity().Identity().(string)] = b

	return b
}

func (br blogRepository) All() []entity.Blog {
	var bs []entity.Blog

	for _, b := range blogs {
		bs = append(bs, b)
	}

	return bs
}

func (br blogRepository) Update(b entity.Blog) (entity.Blog, error) {

	for _, blog := range blogs {
		if blog.Identity().Identity() == b.Identity().Identity() {
			blog.SetName(b.Name())
			return blog, nil
		}
	}
	var nope entity.Blog

	return nope, errors.New("blog not found")
}
