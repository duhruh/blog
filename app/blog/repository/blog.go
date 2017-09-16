package repository

import (
	"github.com/duhruh/blog/app/blog/entity"
	"github.com/duhruh/tackle/domain"
)

type BlogRepository interface {
	FindByIdentity(id domain.Identity) (entity.Blog, error)
	Create(b entity.Blog) entity.Blog
	All() []entity.Blog
}

type blogRepository struct {
}

var blogs map[string]entity.Blog

func NewBlogRepository() BlogRepository {
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
