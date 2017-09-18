package http

var uuidRegex = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}"

type routes struct {
	Method   string
	Path     string
	Endpoint string
	Encoder  string
}

func getRoutes() []routes {
	return []routes{
		{
			"GET",
			"/blogs",
			"ListBlogsEndpoint",
			"ListBlogsEncoder",
		},
		{
			"POST",
			"/blogs",
			"CreateBlogEndpoint",
			"CreateBlogEncoder",
		},
		{
			"PUT",
			"/blogs/{id:" + uuidRegex + "}",
			"UpdateBlogEndpoint",
			"UpdateBlogEncoder",
		},
		{
			"GET",
			"/blogs/{id:" + uuidRegex + "}",
			"ShowBlogEndpoint",
			"ShowBlogEncoder",
		},

		{
			"GET",
			"/blogs/{id:" + uuidRegex + "}/posts",
			"ListPostsEndpoint",
			"ListPostsEncoder",
		},
		{
			"POST",
			"/blogs/{id:" + uuidRegex + "}/posts",
			"CreatePostEndpoint",
			"CreatePostEncoder",
		},
		{
			"GET",
			"/posts/{id:" + uuidRegex + "}",
			"ShowPostEndpoint",
			"ShowPostEncoder",
		},
	}
}
