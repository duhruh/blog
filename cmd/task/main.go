package main

import (
	"os"

	"github.com/duhruh/tackle/task"

	"github.com/duhruh/blog/tasks"
)

func main() {
	runner := task.NewRunner(os.Stdout)

	runner.Register(tasks.NewServerTask())
	runner.Register(tasks.NewGenerateServiceTask())
	runner.Register(tasks.NewGenerateProtoBlogTask())
	runner.Register(tasks.NewRoutesTask())
	runner.Register(tasks.NewBuildTask())

	runner.Run(os.Args)
}
