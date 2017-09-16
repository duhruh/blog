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

func (s *loggingService) AllBlogs() (bs []entity.Blog, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "AllBlogs",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.AllBlogs()
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
