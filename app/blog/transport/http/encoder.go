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
}

func NewEncoderFactory() tacklehttp.EncoderFactory {
	return encoderFactory{
		EncoderFactory: tacklehttp.NewEncoderFactory(),
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
	return kithttp.EncodeResponseFunc(func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		//data := response.(tackle.Packet)
		//err := hs.errorFromResponse(data)
		//if err != nil {
		//	hs.ErrorEncoder()(ctx, err, w)
		//	return nil
		//}
		blog := response.(entity.Blog)

		bjson := struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		}{
			Id:   blog.Identity().Identity().(string),
			Name: blog.Name(),
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return json.NewEncoder(w).Encode(bjson)
	})
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
	return kithttp.EncodeResponseFunc(func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		//data := response.(tackle.Packet)
		//err := hs.errorFromResponse(data)
		//if err != nil {
		//	hs.ErrorEncoder()(ctx, err, w)
		//	return nil
		//}

		blog := response.(entity.Blog)

		bjson := struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		}{
			Id:   blog.Identity().Identity().(string),
			Name: blog.Name(),
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return json.NewEncoder(w).Encode(bjson)
	})
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
	return kithttp.EncodeResponseFunc(func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		data := response.(tackle.Packet)
		err := data.Get("error")
		if err != nil {
			hs.ErrorEncoder()(ctx, err.(error), w)
			return nil
		}
		//err := hs.errorFromResponse(data)
		//if err != nil {
		//
		//}

		//var blogs []entity.Blog

		bb := data.Get("data")
		blogs := bb.([]entity.Blog)

		type bjson struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		}
		var b []bjson

		for _, blog := range blogs {
			b = append(b, bjson{Id: blog.Identity().Identity().(string), Name: blog.Name()})
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return json.NewEncoder(w).Encode(b)
	})
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
	return kithttp.EncodeResponseFunc(func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		//data := response.(tackle.Packet)
		//err := hs.errorFromResponse(data)
		//if err != nil {
		//	hs.ErrorEncoder()(ctx, err, w)
		//	return nil
		//}
		post := response.(entity.Post)

		bjson := struct {
			Id   string `json:"id"`
			Body string `json:"body"`
			Blog string `json:"blog_id"`
		}{
			Id:   post.Identity().Identity().(string),
			Body: post.Body(),
			Blog: post.BlogId().Identity().(string),
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return json.NewEncoder(w).Encode(bjson)
	})
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
	return kithttp.EncodeResponseFunc(func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		//data := response.(tackle.Packet)
		//err := hs.errorFromResponse(data)
		//if err != nil {
		//	hs.ErrorEncoder()(ctx, err, w)
		//	return nil
		//}

		post := response.(entity.Post)

		bjson := struct {
			Id   string `json:"id"`
			Body string `json:"body"`
			Blog string `json:"blog_id"`
		}{
			Id:   post.Identity().Identity().(string),
			Body: post.Body(),
			Blog: post.BlogId().Identity().(string),
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return json.NewEncoder(w).Encode(bjson)
	})
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
	return kithttp.EncodeResponseFunc(func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		//data := response.(tackle.Packet)
		//err := hs.errorFromResponse(data)
		//if err != nil {
		//	hs.ErrorEncoder()(ctx, err, w)
		//	return nil
		//}

		//var blogs []entity.Blog

		posts := response.([]entity.Post)

		type bjson struct {
			Id   string `json:"id"`
			Body string `json:"body"`
			Blog string `json:"blog_id"`
		}
		var b []bjson

		for _, post := range posts {
			b = append(b, bjson{Id: post.Identity().Identity().(string), Body: post.Body(), Blog: post.BlogId().Identity().(string)})
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return json.NewEncoder(w).Encode(b)
	})
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
	return kithttp.EncodeResponseFunc(func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		//data := response.(tackle.Packet)
		//err := hs.errorFromResponse(data)
		//if err != nil {
		//	hs.ErrorEncoder()(ctx, err, w)
		//	return nil
		//}

		//var blogs []entity.Blog

		blog := response.(entity.Blog)

		bjson := struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		}{
			Id:   blog.Identity().Identity().(string),
			Name: blog.Name(),
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return json.NewEncoder(w).Encode(bjson)
	})
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
