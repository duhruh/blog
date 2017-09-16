package tasks

import (
	"io"

	"github.com/duhruh/tackle/task"
)

type HelloTask struct {
	task.Helpers
	shortDescription string
	description      string
	name             string
	options          []task.Option
	arguments        []task.Argument
}

func NewHelloTask() task.Task {
	return HelloTask{
		Helpers:          task.NewHelpers(),
		name:             "hello",
		shortDescription: "Short description here",
		description:      "Description here",
		options:          []task.Option{},
		arguments:        []task.Argument{},
	}
}

func (t HelloTask) ShortDescription() string   { return t.shortDescription }
func (t HelloTask) Description() string        { return t.description }
func (t HelloTask) Name() string               { return t.name }
func (t HelloTask) Options() []task.Option     { return t.options }
func (t HelloTask) Arguments() []task.Argument { return t.arguments }

func (t HelloTask) Run(w io.Writer) {
	t.Say(w, "hello")
}
