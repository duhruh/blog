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
	runner.Register(tasks.NewGenerateProtoBlogTask())
	runner.Register(tasks.NewRoutesTask())

	runner.Run(os.Args)
}
