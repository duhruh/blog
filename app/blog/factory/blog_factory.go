package factory

import "github.com/duhruh/blog/app/blog/entity"

type BlogFactory interface {
	BlogFromImmutable(entity.ImmutableBlog) entity.Blog
}

type blogFactory struct {
}


func NewBlogFactory() BlogFactory{
	return blogFactory{}
}

func (b blogFactory) BlogFromImmutable(u entity.ImmutableBlog) entity.Blog {
	blog := entity.NewBlog()
	blog.SetIdentity(u.Identity())
	blog.SetName(u.Name())
	return blog
}
