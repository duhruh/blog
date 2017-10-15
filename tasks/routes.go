package tasks

import (
	"bytes"
	"context"
	"io"

	"github.com/duhruh/blog/app"

	"github.com/duhruh/tackle"
	"github.com/duhruh/tackle/task"
	"github.com/duhruh/tackle/transport/grpc"
	"github.com/duhruh/tackle/transport/http"
	"github.com/go-kit/kit/log"
)

type RoutesTask struct {
	task.Helpers
	shortDescription string
	description      string
	name             string
	options          []task.Option
	arguments        []task.Argument
}

func NewRoutesTask() task.Task {
	return RoutesTask{
		Helpers:          task.NewHelpers(),
		name:             "routes",
		shortDescription: "List all routes registered",
		description:      "List all routes registered",
		options:          []task.Option{},
		arguments:        []task.Argument{},
	}
}

func (t RoutesTask) ShortDescription() string   { return t.shortDescription }
func (t RoutesTask) Description() string        { return t.description }
func (t RoutesTask) Name() string               { return t.name }
func (t RoutesTask) Options() []task.Option     { return t.options }
func (t RoutesTask) Arguments() []task.Argument { return t.arguments }

func (t RoutesTask) Run(w io.Writer) {
	a := app.NewApplication(context.Background(), app.NewConfigFromYamlFile(tackle.Test, "config/app.yml"), log.NewNopLogger())

	a.Build()

	var buf bytes.Buffer
	buf.WriteString("HTTP Server\n")
	for _, transport := range a.HttpTransport().Transports() {
		t.explainHttpTransport(transport, &buf)
	}

	buf.WriteString("GRPC Server\n")
	for _, transport := range a.GrpcTransport().Transports() {
		t.explainGrpcTransport(transport, &buf)
	}

	t.Say(w, buf.String())
}

func (t RoutesTask) explainHttpTransport(transport http.HttpTransport, buf *bytes.Buffer) {
	routes := transport.Routes()
	for _, route := range routes {
		buf.WriteString("\t[" + route.Method() + "]\t" + route.Path() + "\t-> " + route.Endpoint() + "\n")
	}
}

func (t RoutesTask) explainGrpcTransport(transport grpc.GrpcTransport, buf *bytes.Buffer) {
	handlers := transport.Handlers()
	for _, handle := range handlers {
		buf.WriteString("\t[" + handle.Name() + "]\t-> " + handle.Endpoint() + "\n")
	}
}
