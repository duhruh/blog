package http

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	tacklehttp "github.com/duhruh/tackle/transport/http"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net"
	"sync"
)

type appHttpTransport struct {
	logger     log.Logger
	addr       string
	transports []tacklehttp.HttpTransport
}

func NewHttpTransport(l log.Logger, addr string) tacklehttp.AppHttpTransport {
	return &appHttpTransport{logger: l, addr: addr}
}

func (ht *appHttpTransport) Build(transports []tacklehttp.HttpTransport) {
	ht.transports = transports
	mux := http.NewServeMux()

	//buf := ht.serverStartMessage()
	for _, transport := range ht.transports {
		//ht.explainTransport(transport, &buf)
		transport.NewHandler(mux)
	}

	http.Handle("/", ht.accessControl(mux))
	http.Handle("/metrics", promhttp.Handler())
}

func (ht *appHttpTransport) Start(wg *sync.WaitGroup) {
	listener, err := net.Listen("tcp", ht.addr)
	if err != nil {
		level.Error(ht.logger).Log("transport", "http", "err", err)
	}

	errs := make(chan error, 2)

	wg.Add(3)
	go ht.listen(errs, listener, wg)
	go ht.osSignals(listener, errs, wg)
	go ht.serverClose(errs, wg)
}

func (ht *appHttpTransport) Transports() []tacklehttp.HttpTransport {
	return ht.transports
}

func (ht *appHttpTransport) listen(errs chan error, listener net.Listener, wg *sync.WaitGroup) {
	defer wg.Done()
	level.Info(ht.logger).Log("transport", "http", "address", ht.addr, "message", "listening")
	errs <- http.Serve(listener, nil)
}

func (ht *appHttpTransport) osSignals(listener net.Listener, errs chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	errs <- fmt.Errorf("%s", <-c)
	err := listener.Close()
	if err != nil {
		ht.logger.Log("terminated", <-errs)
	}
}

func (ht *appHttpTransport) serverClose(errs chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	level.Info(ht.logger).Log("terminated", <-errs)
}

func (ht *appHttpTransport) accessControl(h http.Handler) http.Handler {
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
