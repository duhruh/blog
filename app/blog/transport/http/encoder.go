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
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return json.NewEncoder(w).Encode("")
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
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return json.NewEncoder(w).Encode("")
	})
}

func (hs encoderFactory) AllBlogsEncoder() tacklehttp.Encoder {
	return tacklehttp.NewEncoder(hs.allBlogsRequest(), hs.allBlogsResponse())
}

func (hs encoderFactory) allBlogsRequest() kithttp.DecodeRequestFunc {
	return kithttp.DecodeRequestFunc(func(_ context.Context, r *http.Request) (interface{}, error) {

		return tackle.NewPacket(), nil
	})
}

func (hs encoderFactory) allBlogsResponse() kithttp.EncodeResponseFunc {
	return kithttp.EncodeResponseFunc(func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		//data := response.(tackle.Packet)
		//err := hs.errorFromResponse(data)
		//if err != nil {
		//	hs.ErrorEncoder()(ctx, err, w)
		//	return nil
		//}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return json.NewEncoder(w).Encode("")
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
