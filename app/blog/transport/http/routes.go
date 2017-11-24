package http

import (
	"fmt"

	tacklehttp "github.com/duhruh/tackle/transport/http"
)

var uuidRegex = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}"

func getRoutes() []tacklehttp.Route {
	return []tacklehttp.Route{
		tacklehttp.NewRoute(
			"GET",
			"/blogs",
			"ListBlogsEndpoint",
			"ListBlogsEncoder",
		),
		tacklehttp.NewRoute(
			"POST",
			"/blogs",
			"CreateBlogEndpoint",
			"CreateBlogEncoder",
		),
		tacklehttp.NewRoute(
			"PUT",
			fmt.Sprintf("/blogs/{id:%s}", uuidRegex),
			"UpdateBlogEndpoint",
			"UpdateBlogEncoder",
		),
		tacklehttp.NewRoute(
			"GET",
			fmt.Sprintf("/blogs/{id:%s}", uuidRegex),
			"ShowBlogEndpoint",
			"ShowBlogEncoder",
		),

		tacklehttp.NewRoute(
			"GET",
			fmt.Sprintf("/blogs/{id:%s}/posts", uuidRegex),
			"ListPostsEndpoint",
			"ListPostsEncoder",
		),
		tacklehttp.NewRoute(
			"POST",
			fmt.Sprintf("/blogs/{id:%s}/posts", uuidRegex),
			"CreatePostEndpoint",
			"CreatePostEncoder",
		),
		tacklehttp.NewRoute(
			"GET",
			fmt.Sprintf("/posts/{id:%s}", uuidRegex),
			"ShowPostEndpoint",
			"ShowPostEncoder",
		),
	}
}
