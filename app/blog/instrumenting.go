package blog

import (
	"github.com/duhruh/blog/app/blog/entity"
	"github.com/duhruh/tackle/domain"
	"github.com/go-kit/kit/metrics"
	"time"
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

func (s *instrumentingService) AllBlogs() (bs []entity.Blog, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "AllBlogs").Add(1)
		s.requestLatency.With("method", "AllBlogs").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.Service.AllBlogs()
}

func (s *instrumentingService) CreateBlog(name string) (bs entity.Blog, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "CreateBlog").Add(1)
		s.requestLatency.With("method", "CreateBlog").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.Service.CreateBlog(name)
}
