package blog

import (
	"errors"

	"github.com/duhruh/blog/app/blog/entity"
	"github.com/duhruh/blog/app/blog/repository"

	"github.com/duhruh/tackle/domain"
)

var ErrBlogNotFound = errors.New("blog not found")
var ErrPostNotFound = errors.New("post not found")

type Service interface {
	ShowBlog(id domain.Identity) (entity.Blog, error)
	ListBlogs() ([]entity.Blog, error)
	CreateBlog(name string) (entity.Blog, error)
	UpdateBlog(blog entity.Blog) (entity.Blog, error)

	ShowPost(id domain.Identity) (entity.Post, error)
	ListPosts(blog entity.Blog) ([]entity.Post, error)
	CreatePost(blog entity.Blog, body string) (entity.Post, error)
}

type service struct {
	blogRepository repository.BlogRepository
	postRepository repository.PostRepository
}

func newService(blogRepo repository.BlogRepository, postRepo repository.PostRepository) service {
	return service{
		blogRepository: blogRepo,
		postRepository: postRepo,
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

func (s service) ListBlogs() ([]entity.Blog, error) {
	return s.blogRepository.All(), nil
}

func (s service) CreateBlog(name string) (entity.Blog, error) {

	blog := entity.NewBlog()
	blog.SetName(name)
	blog.SetIdentity(entity.NextIdentity())

	return s.blogRepository.Create(blog), nil
}

func (s service) UpdateBlog(blog entity.Blog) (entity.Blog, error) {
	return s.blogRepository.Update(blog)
}

func (s service) ShowPost(id domain.Identity) (entity.Post, error) {
	var p entity.Post

	p, err := s.postRepository.FindByIdentity(id)

	if err != nil {
		return p, ErrPostNotFound
	}

	return p, nil
}

func (s service) ListPosts(blog entity.Blog) ([]entity.Post, error) {
	return s.postRepository.All(), nil
}

func (s service) CreatePost(blog entity.Blog, body string) (entity.Post, error) {
	post := entity.NewPost()
	post.SetIdentity(entity.NextIdentity())
	post.SetBody(body)
	post.SetBlog(blog)

	return s.postRepository.Create(post), nil
}
