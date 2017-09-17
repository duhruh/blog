package http

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	http2 "github.com/duhruh/tackle/transport/http"
	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"sync"
)

type HttpTransport interface {
	Mount(transports []http2.HttpTransport, wg *sync.WaitGroup)
}

type appHttpTransport struct {
	logger log.Logger
	addr   string
}

func NewHttpTransport(l log.Logger, addr string) HttpTransport {
	return appHttpTransport{logger: l, addr: addr}
}

func (ht appHttpTransport) Mount(transports []http2.HttpTransport, wg *sync.WaitGroup) {
	mux := http.NewServeMux()

	for _, transport := range transports {
		transport.NewHandler(mux)
	}

	http.Handle("/", ht.accessControl(mux))
	http.Handle("/metrics", promhttp.Handler())

	errs := make(chan error, 2)

	wg.Add(3)
	go ht.listen(errs, wg)
	go ht.osSignals(errs, wg)
	go ht.serverClose(errs, wg)
}

func (ht appHttpTransport) listen(errs chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	ht.logger.Log("transport", "http", "address", ht.addr, "msg", "listening")
	errs <- http.ListenAndServe(ht.addr, nil)
}

func (ht appHttpTransport) osSignals(errs chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	errs <- fmt.Errorf("%s", <-c)
}

func (ht appHttpTransport) serverClose(errs chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	ht.logger.Log("terminated", <-errs)
}

func (ht appHttpTransport) accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
