package repository

import (
	"github.com/duhruh/blog/app/blog/entity"
	"github.com/duhruh/tackle/domain"
)

type PostRepository interface {
	FindByIdentity(id domain.Identity) (entity.Post, error)
	Create(b entity.Post) entity.Post
	All() []entity.Post
}

type postRepository struct {
}

var posts map[string]entity.Post

func NewPostRepository() PostRepository {
	posts = make(map[string]entity.Post)
	return postRepository{}
}

func (br postRepository) FindByIdentity(id domain.Identity) (entity.Post, error) {
	return posts[id.Identity().(string)], nil
}

func (br postRepository) Create(b entity.Post) entity.Post {
	posts[b.Identity().Identity().(string)] = b

	return b
}

func (br postRepository) All() []entity.Post {
	var bs []entity.Post

	for _, b := range posts {
		bs = append(bs, b)
	}

	return bs
}
