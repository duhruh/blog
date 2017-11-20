package tasks

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/duhruh/tackle"
	"github.com/duhruh/tackle/task"

	"github.com/duhruh/blog/app"
)

const (
	BuildTimeDateFormat = "20060102150405"
	configPath          = "config/app.yml"
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
	config := app.NewConfigFromYamlFile(tackle.Development, configPath)
	build, err := t.GetOption(t.options, "build")
	dir, err := t.GetArgument(t.arguments, "output")
	if err != nil {
		panic(err)
	}

	var (
		outBinName  = fmt.Sprintf("%s/%s_%s", dir.Value(), time.Now().UTC().Format(BuildTimeDateFormat), config.Name())
		gc, _       = exec.Command("git", "rev-parse", "--short", "HEAD").Output()
		buildNumber = build.Value().(string)
		version     = config.Version()
		buildTime   = time.Now().UTC().Format(time.RFC3339Nano)
		cfgPkg      = config.ConfigPath()
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
		config.EntryPoint(),
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
