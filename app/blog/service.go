package blog

import (
	"errors"
	"github.com/duhruh/blog/app/blog/entity"
	"github.com/duhruh/blog/app/blog/repository"
	"github.com/duhruh/tackle/domain"
)

var ErrBlogNotFound = errors.New("blog not found")

type Service interface {
	ShowBlog(id domain.Identity) (entity.Blog, error)
	AllBlogs() ([]entity.Blog, error)
	CreateBlog(name string) (entity.Blog, error)
}

type service struct {
	blogRepository repository.BlogRepository
}

func newService(blogRepo repository.BlogRepository) service {
	return service{
		blogRepository: blogRepo,
	}
}

func (s service) ShowBlog(id domain.Identity) (entity.Blog, error) {
	var b entity.Blog

	b, err := s.blogRepository.FindByIdentity(id)

	if err != nil {
		return b, ErrBlogNotFound
	}

	return b, nil
}

func (s service) AllBlogs() ([]entity.Blog, error) {
	return s.blogRepository.All(), nil
}

func (s service) CreateBlog(name string) (entity.Blog, error) {

	blog := entity.NewBlog()
	blog.SetName(name)

	return s.blogRepository.Create(blog), nil
}
