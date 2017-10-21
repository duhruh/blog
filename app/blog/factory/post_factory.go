package factory

import "github.com/duhruh/blog/app/blog/entity"

type PostFactory interface {
	PostFromImmutable(post entity.ImmutablePost) entity.Post
}

type postFactory struct {
}

func NewPostFactory() PostFactory {
	return postFactory{}
}

func (b postFactory) PostFromImmutable(u entity.ImmutablePost) entity.Post {
	post := entity.NewPost()
	post.SetIdentity(u.Identity())
	post.SetBody(u.Body())
	var bb entity.Blog
	bb.SetIdentity(post.BlogId())
	post.SetBlog(bb)
	return post
}
