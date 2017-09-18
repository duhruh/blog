package blog

import (
	"github.com/duhruh/blog/app/blog/entity"
	"github.com/duhruh/tackle/domain"
	"github.com/go-kit/kit/log"
	"time"
)

type loggingService struct {
	logger log.Logger
	Service
}

// NewLoggingService returns a new instance of a logging Service.
func newLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) ShowBlog(id domain.Identity) (blog entity.Blog, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "ShowBlog",
			"id", id.Identity(),
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.ShowBlog(id)
}

func (s *loggingService) ListBlogs() (bs []entity.Blog, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "ListBlogs",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.ListBlogs()
}

func (s *loggingService) CreateBlog(name string) (bs entity.Blog, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "CreateBlog",
			"name", name,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.CreateBlog(name)
}

func (s *loggingService) ShowPost(id domain.Identity) (post entity.Post, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "ShowPost",
			"id", id.Identity(),
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.ShowPost(id)
}

func (s *loggingService) ListPosts(blog entity.Blog) (bs []entity.Post, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "ListPosts",
			"blog", blog.Identity().Identity(),
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.ListPosts(blog)
}

func (s *loggingService) CreatePost(blog entity.Blog, body string) (bs entity.Post, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "CreatePost",
			"blog", blog.Identity().Identity(),
			"body", body,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.CreatePost(blog, body)
}

func (s *loggingService) UpdateBlog(blog entity.Blog) (_ entity.Blog, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "UpdateBlog",
			"blog", blog.Identity().Identity(),
			"name", blog.Name(),
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.UpdateBlog(blog)
}
