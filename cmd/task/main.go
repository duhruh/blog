package main

import (
	"os"

	"github.com/duhruh/blog/tasks"
	"github.com/duhruh/tackle/task"
)

func main() {
	runner := task.NewRunner(os.Stdout)

	runner.Register(tasks.NewServerTask())
	runner.Register(tasks.NewGenerateServiceTask())

	runner.Run(os.Args)
}
