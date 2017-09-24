package grpc

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	tacklegrpc "github.com/duhruh/tackle/transport/grpc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
)

type appGrpcTransport struct {
	logger     log.Logger
	addr       string
	baseServer *grpc.Server
	transports []tacklegrpc.GrpcTransport
}

func NewGrpcTransport(l log.Logger, addr string) tacklegrpc.AppGrpcTransport {
	return &appGrpcTransport{logger: l, addr: addr}
}

func (gt *appGrpcTransport) Build(transports []tacklegrpc.GrpcTransport) {
	gt.baseServer = grpc.NewServer()
	gt.transports = transports

	for _, transport := range gt.transports {
		transport.NewHandler(gt.baseServer)
	}
}

func (gt *appGrpcTransport) Start(wg *sync.WaitGroup) {
	grpcListener, err := net.Listen("tcp", gt.addr)
	if err != nil {
		level.Error(gt.logger).Log("transport", "grpc", "err", err)
	}

	wg.Add(3)
	errs := make(chan error, 2)
	go gt.listen(grpcListener, errs, wg)
	go gt.osSignals(grpcListener, errs, wg)
	go gt.serverClose(errs, wg)
}

func (gt *appGrpcTransport) Transports() []tacklegrpc.GrpcTransport {
	return gt.transports
}

func (gt *appGrpcTransport) listen(grpcListener net.Listener, errs chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	level.Info(gt.logger).Log("transport", "grpc", "addr", gt.addr, "message", "listening")
	errs <- gt.baseServer.Serve(grpcListener)
}

func (gt *appGrpcTransport) osSignals(grpcListener net.Listener, errs chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	errs <- fmt.Errorf("%s", <-c)
	err := grpcListener.Close()
	if err != nil {
		level.Info(gt.logger).Log("terminated", <-errs)
	}
}

func (gt *appGrpcTransport) serverClose(errs chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	level.Info(gt.logger).Log("terminated", <-errs)
}
