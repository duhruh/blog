package tasks

import (
	"io"

	"github.com/duhruh/tackle/task"
	"os"
	"os/exec"
	"time"
)

type ServerTask struct {
	task.Helpers
	shortDescription string
	description      string
	name             string
	options          []task.Option
	arguments        []task.Argument
}

func NewServerTask() task.Task {
	return ServerTask{
		Helpers:          task.NewHelpers(),
		name:             "server",
		shortDescription: "Runs the tackle server",
		description:      "Runs the tackle server",
		options:          []task.Option{},
		arguments:        []task.Argument{},
	}
}

func (t ServerTask) ShortDescription() string   { return t.shortDescription }
func (t ServerTask) Description() string        { return t.description }
func (t ServerTask) Name() string               { return t.name }
func (t ServerTask) Options() []task.Option     { return t.options }
func (t ServerTask) Arguments() []task.Argument { return t.arguments }

func (t ServerTask) Run(w io.Writer) {
	var (
		gitCommit   = "ec71764"
		buildNumber = "1"
		version     = "v1.0.0"
		buildTime   = time.Now().UTC().Format(time.RFC3339Nano)
		cfgPkg      = "github.com/duhruh/blog/config"
	)

	cmd := exec.Command(
		"go",
		"run",
		"-ldflags",
		""+t.ldflag(cfgPkg, "GitCommit", gitCommit)+" "+
			t.ldflag(cfgPkg, "BuildNumber", buildNumber)+" "+
			t.ldflag(cfgPkg, "Version", version)+" "+
			t.ldflag(cfgPkg, "BuildTime", buildTime),
		"cmd/api/main.go",
		"-http-bind-address=:8081",
	)
	cmd.Env = os.Environ()
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func (t ServerTask) ldflag(pkg string, variable string, value string) string {
	return "-X \"" + pkg + "." + variable + "=" + value + "\""
}
