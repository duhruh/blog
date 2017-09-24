package grpc

import (
	"fmt"
	tacklegrpc "github.com/duhruh/tackle/transport/grpc"
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type GrpcTransport interface {
	Mount(transports []tacklegrpc.GrpcTransport, wg *sync.WaitGroup)
}

type appGrpcTransport struct {
	logger log.Logger
	addr   string
}

func NewGrpcTransport(l log.Logger, addr string) GrpcTransport {
	return appGrpcTransport{logger: l, addr: addr}
}

func (gt appGrpcTransport) Mount(transports []tacklegrpc.GrpcTransport, wg *sync.WaitGroup) {
	baseServer := grpc.NewServer()

	grpcListener, err := net.Listen("tcp", gt.addr)
	if err != nil {
		gt.logger.Log("transport", "grpc", "err", err)
	}

	for _, transport := range transports {
		transport.NewHandler(baseServer)
	}

	wg.Add(3)
	errs := make(chan error, 2)
	go gt.listen(baseServer, grpcListener, errs, wg)
	go gt.osSignals(grpcListener, errs, wg)
	go gt.serverClose(errs, wg)

}

func (gt appGrpcTransport) listen(baseServer *grpc.Server, grpcListener net.Listener, errs chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	gt.logger.Log("transport", "grpc", "addr", gt.addr, "msg", "listening")
	errs <- baseServer.Serve(grpcListener)
}

func (gt appGrpcTransport) osSignals(grpcListener net.Listener, errs chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	errs <- fmt.Errorf("%s", <-c)
	err := grpcListener.Close()
	if err != nil {
		gt.logger.Log("terminated", <-errs)
	}
}

func (gt appGrpcTransport) serverClose(errs chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	gt.logger.Log("terminated", <-errs)
}
