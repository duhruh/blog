package tasks

import (
	"io"

	"github.com/duhruh/tackle/task"
	"os"
	"os/exec"
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
	exec.Command("git rev-parse ")
	cmd := exec.Command(
		"go",
		"run",
		"-ldflags",
		"-X \"github.com/duhruh/blog/config.GitCommit=82847ca\" -X \"github.com/duhruh/blog/config.Version=v1.0.0\" -X \"github.com/duhruh/scaffold/config.BuildTime=Sat, Sep  9, 2017  8:01:22 PM\"",
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
