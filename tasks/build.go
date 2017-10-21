package tasks

import (
	"fmt"
	"github.com/duhruh/tackle/task"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

type BuildTask struct {
	task.Helpers
	shortDescription string
	description      string
	name             string
	options          []task.Option
	arguments        []task.Argument
}

func NewBuildTask() task.Task {
	return BuildTask{
		Helpers:          task.NewHelpers(),
		name:             "build",
		shortDescription: "Builds the project",
		description:      "Builds the project",
		options: []task.Option{
			task.NewOption("build", "the build number"),
			task.NewOption("version", "the version"),
		},
		arguments: []task.Argument{
			task.NewArgument("output", "the output directory"),
		},
	}
}

func (t BuildTask) ShortDescription() string   { return t.shortDescription }
func (t BuildTask) Description() string        { return t.description }
func (t BuildTask) Name() string               { return t.name }
func (t BuildTask) Options() []task.Option     { return t.options }
func (t BuildTask) Arguments() []task.Argument { return t.arguments }

func (t BuildTask) Run(w io.Writer) {

	build, err := t.GetOption(t.options, "build")
	ve, err := t.GetOption(t.options, "version")
	dir, err := t.GetArgument(t.arguments, "output")
	if err != nil {
		panic(err)
	}

	var (
		outBinName  = fmt.Sprintf("%s/%s_blog", dir.Value(), time.Now().UTC().Format("20060102"))
		gc, _       = exec.Command("git", "rev-parse", "--short", "HEAD").Output()
		buildNumber = build.Value().(string)
		version     = ve.Value().(string)
		buildTime   = time.Now().UTC().Format(time.RFC3339Nano)
		cfgPkg      = "github.com/duhruh/blog/config"
		gitCommit   = strings.Trim(string(gc), "\n")
	)

	cmd := exec.Command(
		"go",
		"build",
		"-o",
		outBinName,
		"-ldflags",
		""+t.ldflag(cfgPkg, "GitCommit", gitCommit)+" "+
			t.ldflag(cfgPkg, "BuildNumber", buildNumber)+" "+
			t.ldflag(cfgPkg, "Version", version)+" "+
			t.ldflag(cfgPkg, "BuildTime", buildTime),
		"cmd/api/main.go",
	)
	cmd.Env = os.Environ()
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
func (t BuildTask) ldflag(pkg string, variable string, value string) string {
	return "-X \"" + pkg + "." + variable + "=" + value + "\""
}
