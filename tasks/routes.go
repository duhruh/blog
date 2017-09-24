package tasks

import (
	"io"
	"bytes"

	"github.com/duhruh/tackle/task"
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
		shortDescription: "List the http routes registered",
		description:      "List the http routes registered",
		options: []task.Option{},
		arguments: []task.Argument{},
	}
}

func (t RoutesTask) ShortDescription() string   { return t.shortDescription }
func (t RoutesTask) Description() string        { return t.description }
func (t RoutesTask) Name() string               { return t.name }
func (t RoutesTask) Options() []task.Option     { return t.options }
func (t RoutesTask) Arguments() []task.Argument { return t.arguments }

func (t RoutesTask) Run(w io.Writer) {

}
