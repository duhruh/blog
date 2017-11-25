package tasks

import (
	"io"
	"os"
	"path/filepath"
	"text/template"

	"github.com/duhruh/tackle/task"
)

type GenerateServiceTask struct {
	task.Helpers
	shortDescription string
	description      string
	name             string
	options          []task.Option
	arguments        []task.Argument
	writer           io.Writer
}

func NewGenerateServiceTask() task.Task {
	return GenerateServiceTask{
		Helpers:          task.NewHelpers(),
		name:             "generate:service",
		shortDescription: "Generates a service fully instrumented",
		description:      "Generates a service with logging, instrumentation, and endpoint",
		options:          []task.Option{},
		arguments: []task.Argument{
			task.NewArgument("name", "the name of the service to generate"),
		},
	}
}

func (t GenerateServiceTask) ShortDescription() string   { return t.shortDescription }
func (t GenerateServiceTask) Description() string        { return t.description }
func (t GenerateServiceTask) Name() string               { return t.name }
func (t GenerateServiceTask) Options() []task.Option     { return t.options }
func (t GenerateServiceTask) Arguments() []task.Argument { return t.arguments }

func (t GenerateServiceTask) Run(w io.Writer) {
	t.writer = w
	var newService servicePackage

	taskArg, err := t.GetArgument(t.arguments, "name")
	if err != nil {
		panic(err)
	}

	newService.Package = taskArg.Value().(string)

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	d := filepath.Join(dir, "app", newService.Package)

	err = os.Mkdir(d, 0755)
	if err != nil {
		panic(err)
	}

	t.generateFileWithTemplate(newService, "service", t.serviceTemplate())
	t.generateFileWithTemplate(newService, "instrumenting", t.instrumentingTemplate())
	t.generateFileWithTemplate(newService, "endpoint", t.endpointTemplate())
	t.generateFileWithTemplate(newService, "logging", t.loggingTemplate())

	t.Say(w, "All done!")
}

type servicePackage struct {
	Package string
}

func (t GenerateServiceTask) generateFileWithTemplate(pkg servicePackage, fileName string, temp string) {
	tmpl, err := template.New(pkg.Package + "_" + fileName).Parse(temp)
	if err != nil {
		panic(err)
	}

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fullPath := filepath.Join(dir, "app", pkg.Package, fileName+".go")
	t.Say(t.writer, "generating file: "+fullPath)

	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(file, pkg)
	if err != nil {
		panic(err)
	}
}

func (t GenerateServiceTask) serviceTemplate() string {
	return `package {{.Package}}


type Service interface {
	/*
    Foo(arg string) (string, error)
	*/
}

type service struct{}

func newService() *service {
	return &service{}
}

/*
func (s *service) Foo(arg string) (string, error) {
	return "example", nil
}
*/

`
}
func (t GenerateServiceTask) loggingTemplate() string {
	return `package {{.Package}}

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	Service
}

func newLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}
/*
func (s *loggingService) Foo(arg string) (arg string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "Foo",
			"arg", arg,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.Foo(arg)
}
*/

`
}

func (t GenerateServiceTask) instrumentingTemplate() string {
	return `package {{.Package}}

import (
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

func newInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}

/*
func (s *instrumentingService) Foo(arg string) (string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Foo").Add(1)
		s.requestLatency.With("method", "Foo").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Foo(arg)
}
*/

`
}

func (t GenerateServiceTask) endpointTemplate() string {
	return `package {{.Package}}

import (
	"context"

	"github.com/duhruh/tackle"
	"github.com/go-kit/kit/endpoint"
)

type endpointFactory struct {
	tackle.EndpointFactory
	service    Service
	serializer Serializer
}

func newEndpointFactory(s Service, se Serializer) tackle.EndpointFactory {
	return endpointFactory{
		EndpointFactory: tackle.NewEndpointFactory(),
		service:         s,
		serializer:      se,
	}
}

func (ef endpointFactory) Generate(end string) (endpoint.Endpoint, error) {
	return ef.EndpointFactory.GenerateWithInstance(ef, end)
}


`
}
