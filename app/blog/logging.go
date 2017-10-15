package blog

import (
	"time"

	"github.com/duhruh/blog/app/blog/entity"

	"github.com/duhruh/blog/app/blog/errors"
	"github.com/duhruh/tackle/domain"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type loggingService struct {
	logger log.Logger
	Service
}

// NewLoggingService returns a new instance of a logging Service.
func newLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) logIt(took time.Duration, err error, keyvals ...interface{}) {
	var logLevel log.Logger
	logLevel = level.Debug(s.logger)

	if err != nil {
		logLevel = level.Error(s.logger)
	}

	keyvals = append(keyvals, "took", took, "error", err, "trace", errors.StackTrace(err))

	logLevel.Log(keyvals...)
}

func (s *loggingService) ShowBlog(id domain.Identity) (blog entity.Blog, err error) {
	defer func(begin time.Time) {
		s.logIt(time.Since(begin), err, "method", "ShowBlog", "id", id.Identity())
	}(time.Now())
	return s.Service.ShowBlog(id)
}

func (s *loggingService) ListBlogs() (bs []entity.Blog, err error) {
	defer func(begin time.Time) {
		s.logIt(time.Since(begin), err, "method", "ListBlogs")
	}(time.Now())
	return s.Service.ListBlogs()
}

func (s *loggingService) CreateBlog(name string) (bs entity.Blog, err error) {
	defer func(begin time.Time) {
		s.logIt(time.Since(begin), err, "method", "CreateBlog", "name", name)
	}(time.Now())
	return s.Service.CreateBlog(name)
}

func (s *loggingService) ShowPost(id domain.Identity) (post entity.Post, err error) {
	defer func(begin time.Time) {
		s.logIt(time.Since(begin), err, "method", "ShowPost", "id", id.Identity())
	}(time.Now())
	return s.Service.ShowPost(id)
}

func (s *loggingService) ListPosts(blog entity.Blog) (bs []entity.Post, err error) {
	defer func(begin time.Time) {
		s.logIt(time.Since(begin), err, "method", "ListPosts", "blog", blog.Identity().Identity())
	}(time.Now())
	return s.Service.ListPosts(blog)
}

func (s *loggingService) CreatePost(blog entity.Blog, body string) (bs entity.Post, err error) {
	defer func(begin time.Time) {
		s.logIt(time.Since(begin), err, "method", "CreatePost", "blog", blog.Identity().Identity(), "body", body)
	}(time.Now())
	return s.Service.CreatePost(blog, body)
}

func (s *loggingService) UpdateBlog(blog entity.Blog) (_ entity.Blog, err error) {
	defer func(begin time.Time) {
		s.logIt(time.Since(begin), err, "method", "UpdateBlog", "blog", blog.Identity().Identity(), "name", blog.Name())
	}(time.Now())
	return s.Service.UpdateBlog(blog)
}
