package grpc

import (
	"fmt"
	tacklegrpc "github.com/duhruh/tackle/transport/grpc"
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type GrpcTransport interface {
	Mount(transports []tacklegrpc.GrpcTransport)
}

type appGrpcTransport struct {
	logger log.Logger
	addr   string
}

func NewGrpcTransport(l log.Logger, addr string) GrpcTransport {
	return appGrpcTransport{logger: l, addr: addr}
}

func (gt appGrpcTransport) Mount(transports []tacklegrpc.GrpcTransport) {
	baseServer := grpc.NewServer()

	grpcListener, err := net.Listen("tcp", gt.addr)
	if err != nil {
		gt.logger.Log("transport", "grpc", "err", err)
		panic(err)
	}

	for _, transport := range transports {
		transport.NewHandler(baseServer)
	}

	errs := make(chan error, 2)
	go func() {
		gt.logger.Log("transport", "grpc", "addr", gt.addr, "msg", "listening")
		errs <- baseServer.Serve(grpcListener)
	}()
	func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
		grpcListener.Close()
	}()

	go func() {
		gt.logger.Log("terminated", <-errs)
	}()

}
