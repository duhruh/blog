package grpc

import (
	tacklegrpc "github.com/duhruh/tackle/transport/grpc"
)

const (
	BlogServiceListBlogs = "BlogService.listBlogs"
)

func getHandlers() []tacklegrpc.Handler {
	return []tacklegrpc.Handler{
		tacklegrpc.NewHandler(
			BlogServiceListBlogs,
			"ListBlogsEndpoint",
			"ListBlogsEncoder",
		),
	}
}
