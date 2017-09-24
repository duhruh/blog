package http

import (
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
			"/blogs/{id:"+uuidRegex+"}",
			"UpdateBlogEndpoint",
			"UpdateBlogEncoder",
		),
		tacklehttp.NewRoute(
			"GET",
			"/blogs/{id:"+uuidRegex+"}",
			"ShowBlogEndpoint",
			"ShowBlogEncoder",
		),

		tacklehttp.NewRoute(
			"GET",
			"/blogs/{id:"+uuidRegex+"}/posts",
			"ListPostsEndpoint",
			"ListPostsEncoder",
		),
		tacklehttp.NewRoute(
			"POST",
			"/blogs/{id:"+uuidRegex+"}/posts",
			"CreatePostEndpoint",
			"CreatePostEncoder",
		),
		tacklehttp.NewRoute(
			"GET",
			"/posts/{id:"+uuidRegex+"}",
			"ShowPostEndpoint",
			"ShowPostEncoder",
		),
	}
}
