package http

import (
	"encoding/json"
	"errors"

	"github.com/duhruh/blog/app/blog/entity"
)

type Serializer interface {
	JsonBlog(b entity.Blog) ([]byte, error)
	JsonBlogs(blogs []entity.Blog) ([]byte, error)

	JsonPost(b entity.Post) ([]byte, error)
	JsonPosts(posts []entity.Post) ([]byte, error)
}

type httpSerializer struct {
}

func NewSerializer() Serializer {
	return httpSerializer{}
}

type jsonBlog struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type jsonPost struct {
	Id   string `json:"id"`
	Body string `json:"body"`
	Blog string `json:"blog_id"`
}

func (hs httpSerializer) JsonBlog(b entity.Blog) ([]byte, error) {
	bytes, err := json.Marshal(jsonBlog{
		Id:   b.Identity().Identity().(string),
		Name: b.Name(),
	})

	if err != nil {
		return []byte{}, errors.New("unable to marshal blog")
	}

	return bytes, nil
}

func (hs httpSerializer) JsonBlogs(blogs []entity.Blog) ([]byte, error) {
	var b []jsonBlog

	for _, blog := range blogs {
		b = append(b, jsonBlog{Id: blog.Identity().Identity().(string), Name: blog.Name()})
	}

	bytes, err := json.Marshal(b)

	if err != nil {
		return []byte{}, errors.New("unable to marshal blogs")
	}

	return bytes, nil
}

func (hs httpSerializer) JsonPost(p entity.Post) ([]byte, error) {
	bytes, err := json.Marshal(jsonPost{
		Id:   p.Identity().Identity().(string),
		Body: p.Body(),
		Blog: p.BlogId().Identity().(string),
	})

	if err != nil {
		return []byte{}, errors.New("unable to marshal post")
	}

	return bytes, nil
}

func (hs httpSerializer) JsonPosts(posts []entity.Post) ([]byte, error) {
	var pp []jsonPost

	for _, post := range posts {
		pp = append(pp, jsonPost{Id: post.Identity().Identity().(string), Body: post.Body(), Blog: post.BlogId().Identity().(string)})
	}

	bytes, err := json.Marshal(pp)

	if err != nil {
		return []byte{}, errors.New("unable to marshal posts")
	}

	return bytes, nil
}
