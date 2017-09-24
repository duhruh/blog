package tasks

import (
	"io"
	"os/exec"

	"github.com/duhruh/tackle/task"
)

type GenerateProtoBlogTask struct {
	task.Helpers
	shortDescription string
	description      string
	name             string
	options          []task.Option
	arguments        []task.Argument
}

func NewGenerateProtoBlogTask() task.Task {
	return GenerateProtoBlogTask{
		Helpers:          task.NewHelpers(),
		name:             "generate:proto:blog",
		shortDescription: "Regenerates the go proto buf file for blog",
		description:      "Regenerates the go proto buf file for blog",
		options:          []task.Option{},
		arguments:        []task.Argument{},
	}
}

func (t GenerateProtoBlogTask) ShortDescription() string   { return t.shortDescription }
func (t GenerateProtoBlogTask) Description() string        { return t.description }
func (t GenerateProtoBlogTask) Name() string               { return t.name }
func (t GenerateProtoBlogTask) Options() []task.Option     { return t.options }
func (t GenerateProtoBlogTask) Arguments() []task.Argument { return t.arguments }

func (t GenerateProtoBlogTask) Run(w io.Writer) {

	cmd := exec.Command("protoc", "app/blog/proto/blog.proto", "--go_out=plugins=grpc:.")

	err := cmd.Run()

	if err != nil {
		panic(err)
	}

	t.Say(w, "All done")
}
