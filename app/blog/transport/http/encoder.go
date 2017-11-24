package http

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/duhruh/tackle"
	tacklehttp "github.com/duhruh/tackle/transport/http"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/duhruh/blog/app/blog/entity"
)

type encoderFactory struct {
	tacklehttp.EncoderFactory

	serializer Serializer
}

func NewEncoderFactory(s Serializer) tacklehttp.EncoderFactory {
	return encoderFactory{
		EncoderFactory: tacklehttp.NewEncoderFactory(),
		serializer:     s,
	}
}

func (ef encoderFactory) Generate(e string) (tacklehttp.Encoder, error) {
	return ef.GenerateWithInstance(ef, e)
}

func (hs encoderFactory) CreateBlogEncoder() tacklehttp.Encoder {
	return tacklehttp.NewEncoder(hs.createBlogRequest(), hs.createBlogResponse())
}

func (hs encoderFactory) createBlogRequest() kithttp.DecodeRequestFunc {
	return kithttp.DecodeRequestFunc(func(_ context.Context, r *http.Request) (interface{}, error) {
		name := r.FormValue("name")
		packet := tackle.NewPacket()
		packet.Put("name", name)
		return packet, nil
	})
}

func (hs encoderFactory) createBlogResponse() kithttp.EncodeResponseFunc {
	return kithttp.EncodeResponseFunc(hs.errorWrap(hs.blogResponse))
}

func (hs encoderFactory) ShowBlogEncoder() tacklehttp.Encoder {
	return tacklehttp.NewEncoder(hs.showBlogRequest(), hs.showBlogResponse())
}

func (hs encoderFactory) showBlogRequest() kithttp.DecodeRequestFunc {
	return kithttp.DecodeRequestFunc(func(_ context.Context, r *http.Request) (interface{}, error) {
		vars := mux.Vars(r)
		packet := tackle.NewPacket()
		packet.Put("id", vars["id"])
		return packet, nil
	})
}

func (hs encoderFactory) showBlogResponse() kithttp.EncodeResponseFunc {
	return kithttp.EncodeResponseFunc(hs.errorWrap(hs.blogResponse))
}

func (hs encoderFactory) ListBlogsEncoder() tacklehttp.Encoder {
	return tacklehttp.NewEncoder(hs.listBlogsRequest(), hs.listBlogsResponse())
}

func (hs encoderFactory) listBlogsRequest() kithttp.DecodeRequestFunc {
	return kithttp.DecodeRequestFunc(func(_ context.Context, r *http.Request) (interface{}, error) {
		return tackle.NewPacket(), nil
	})
}

func (hs encoderFactory) listBlogsResponse() kithttp.EncodeResponseFunc {
	return kithttp.EncodeResponseFunc(hs.errorWrap(hs.blogsResponse))
}

func (hs encoderFactory) CreatePostEncoder() tacklehttp.Encoder {
	return tacklehttp.NewEncoder(hs.createPostRequest(), hs.createPostResponse())
}

func (hs encoderFactory) createPostRequest() kithttp.DecodeRequestFunc {
	return kithttp.DecodeRequestFunc(func(_ context.Context, r *http.Request) (interface{}, error) {
		vars := mux.Vars(r)

		packet := tackle.NewPacket()
		packet.Put("blog_id", vars["id"])

		body := r.FormValue("body")
		packet.Put("body", body)
		return packet, nil
	})
}

func (hs encoderFactory) createPostResponse() kithttp.EncodeResponseFunc {
	return kithttp.EncodeResponseFunc(hs.errorWrap(hs.postResponse))
}

func (hs encoderFactory) ShowPostEncoder() tacklehttp.Encoder {
	return tacklehttp.NewEncoder(hs.showPostRequest(), hs.showPostResponse())
}

func (hs encoderFactory) showPostRequest() kithttp.DecodeRequestFunc {
	return kithttp.DecodeRequestFunc(func(_ context.Context, r *http.Request) (interface{}, error) {
		vars := mux.Vars(r)
		packet := tackle.NewPacket()
		packet.Put("id", vars["id"])
		return packet, nil
	})
}

func (hs encoderFactory) showPostResponse() kithttp.EncodeResponseFunc {
	return kithttp.EncodeResponseFunc(hs.errorWrap(hs.postResponse))
}

func (hs encoderFactory) ListPostsEncoder() tacklehttp.Encoder {
	return tacklehttp.NewEncoder(hs.listPostsRequest(), hs.listPostsResponse())
}

func (hs encoderFactory) listPostsRequest() kithttp.DecodeRequestFunc {
	return kithttp.DecodeRequestFunc(func(_ context.Context, r *http.Request) (interface{}, error) {
		vars := mux.Vars(r)
		packet := tackle.NewPacket()
		packet.Put("blog_id", vars["id"])

		return packet, nil
	})
}

func (hs encoderFactory) listPostsResponse() kithttp.EncodeResponseFunc {
	return kithttp.EncodeResponseFunc(hs.errorWrap(hs.postsResponse))
}

func (hs encoderFactory) UpdateBlogEncoder() tacklehttp.Encoder {
	return tacklehttp.NewEncoder(hs.updateBlogRequest(), hs.updateBlogResponse())
}

func (hs encoderFactory) updateBlogRequest() kithttp.DecodeRequestFunc {
	return kithttp.DecodeRequestFunc(func(_ context.Context, r *http.Request) (interface{}, error) {
		vars := mux.Vars(r)
		packet := tackle.NewPacket()
		packet.Put("id", vars["id"])

		packet.Put("name", r.FormValue("name"))

		return packet, nil
	})
}

func (hs encoderFactory) updateBlogResponse() kithttp.EncodeResponseFunc {
	return kithttp.EncodeResponseFunc(hs.errorWrap(hs.blogResponse))
}

func (hs encoderFactory) ErrorEncoder() kithttp.ErrorEncoder {
	return kithttp.ErrorEncoder(func(_ context.Context, err error, w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		switch err {
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
	})
}

func (hs encoderFactory) errorFromResponse(response interface{}) error {
	err := reflect.ValueOf(response).FieldByName("Err").Interface()
	if e, ok := err.(error); ok && e != nil {
		return e
	}
	return nil
}

func (hs encoderFactory) errorWrap(next tackleResponse) kithttp.EncodeResponseFunc {
	return kithttp.EncodeResponseFunc(func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		pkt := response.(tackle.Packet)
		if hs.handleErrorResponse(ctx, w, pkt) {
			return nil
		}
		return next(ctx, w, pkt)
	})
}

type tackleResponse func(ctx context.Context, w http.ResponseWriter, response tackle.Packet) error

func (hs encoderFactory) blogResponse(ctx context.Context, w http.ResponseWriter, pkt tackle.Packet) error {
	blog := pkt.Get("data").(entity.Blog)
	bytes, err := hs.serializer.JsonBlog(blog)

	return hs.writeJsonResponse(ctx, w, bytes, err)
}

func (hs encoderFactory) blogsResponse(ctx context.Context, w http.ResponseWriter, pkt tackle.Packet) error {
	blogs := pkt.Get("data").([]entity.Blog)
	bytes, err := hs.serializer.JsonBlogs(blogs)

	return hs.writeJsonResponse(ctx, w, bytes, err)
}

func (hs encoderFactory) postsResponse(ctx context.Context, w http.ResponseWriter, pkt tackle.Packet) error {
	posts := pkt.Get("data").([]entity.Post)
	bytes, err := hs.serializer.JsonPosts(posts)

	return hs.writeJsonResponse(ctx, w, bytes, err)
}

func (hs encoderFactory) postResponse(ctx context.Context, w http.ResponseWriter, pkt tackle.Packet) error {
	post := pkt.Get("data").(entity.Post)
	bytes, err := hs.serializer.JsonPost(post)

	return hs.writeJsonResponse(ctx, w, bytes, err)
}

func (hs encoderFactory) handleErrorResponse(ctx context.Context, w http.ResponseWriter, pkt tackle.Packet) bool {
	e := pkt.Get("error")
	if e != nil {
		hs.ErrorEncoder()(ctx, e.(error), w)
		return true
	}

	return false
}

func (hs encoderFactory) writeJsonResponse(ctx context.Context, w http.ResponseWriter, bytes []byte, err error) error {
	if err != nil {
		hs.ErrorEncoder()(ctx, err, w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = w.Write(bytes)
	return err
}
