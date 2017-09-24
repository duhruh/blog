package blog

import (
	"time"

	"github.com/duhruh/blog/app/blog/entity"

	"github.com/duhruh/tackle/domain"
	"github.com/go-kit/kit/metrics"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func newInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}

func (s *instrumentingService) ShowBlog(id domain.Identity) (blog entity.Blog, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "ShowBlog").Add(1)
		s.requestLatency.With("method", "ShowBlog").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.ShowBlog(id)
}

func (s *instrumentingService) ListBlogs() (bs []entity.Blog, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "ListBlogs").Add(1)
		s.requestLatency.With("method", "ListBlogs").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.Service.ListBlogs()
}

func (s *instrumentingService) CreateBlog(name string) (bs entity.Blog, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "CreateBlog").Add(1)
		s.requestLatency.With("method", "CreateBlog").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.Service.CreateBlog(name)
}

func (s *instrumentingService) ShowPost(id domain.Identity) (post entity.Post, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "ShowPost").Add(1)
		s.requestLatency.With("method", "ShowPost").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.ShowPost(id)
}

func (s *instrumentingService) ListPosts(blog entity.Blog) (p []entity.Post, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "ListPosts").Add(1)
		s.requestLatency.With("method", "ListPosts").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.Service.ListPosts(blog)
}

func (s *instrumentingService) CreatePost(blog entity.Blog, body string) (p entity.Post, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "CreatePost").Add(1)
		s.requestLatency.With("method", "CreatePost").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.Service.CreatePost(blog, body)
}

func (s *instrumentingService) UpdateBlog(blog entity.Blog) (_ entity.Blog, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "UpdateBlog").Add(1)
		s.requestLatency.With("method", "UpdateBlog").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.Service.UpdateBlog(blog)
}
