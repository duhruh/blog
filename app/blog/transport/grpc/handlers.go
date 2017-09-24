package grpc

import (
	tacklegrpc "github.com/duhruh/tackle/transport/grpc"
)

func getHandlers() []tacklegrpc.Handler {
	return []tacklegrpc.Handler{
		tacklegrpc.NewHandler(
			"BlogService.listBlogs",
			"ListBlogsEndpoint",
			"ListBlogsEncoder",
		),
	}
}
