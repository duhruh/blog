package tasks

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/duhruh/tackle/dsnotify"
	"github.com/duhruh/tackle/task"
	"github.com/fsnotify/fsnotify"
)

type ServerTask struct {
	task.Helpers
	shortDescription string
	description      string
	name             string
	options          []task.Option
	arguments        []task.Argument

	currentCommand *exec.Cmd
	writer         io.Writer
	outBinName     string
}

func NewServerTask() task.Task {
	return &ServerTask{
		Helpers:          task.NewHelpers(),
		name:             "server",
		shortDescription: "Runs the tackle server",
		description:      "Runs the tackle server",
		options: []task.Option{
			task.NewOption("watch", "recompiles on file change"),
		},
		arguments: []task.Argument{},
	}
}

func (t *ServerTask) ShortDescription() string   { return t.shortDescription }
func (t *ServerTask) Description() string        { return t.description }
func (t *ServerTask) Name() string               { return t.name }
func (t *ServerTask) Options() []task.Option     { return t.options }
func (t *ServerTask) Arguments() []task.Argument { return t.arguments }

func (t *ServerTask) Run(w io.Writer) {
	t.writer = w

	t.outBinName = fmt.Sprintf("/tmp/go-bin-%s", time.Now().String())

	t.Say(w, "\nWelcome to Tackle v1.0.0\n")

	option, _ := t.GetOption(t.options, "watch")
	if option.Value() == nil {
		t.compileAndRun()
		t.currentCommand.Wait()
	}

	directoryWatcher, err := dsnotify.NewDirectoryWatcher()
	if err != nil {
		panic(err)
	}

	directoryWatcher.IgnoreRegex(regexp.MustCompile(`^.glide`))
	directoryWatcher.IgnoreRegex(regexp.MustCompile(`^.git`))
	directoryWatcher.AddDirectory("./")
	directoryWatcher.RegisterFileRegex(regexp.MustCompile(`(.+\.go)`))

	defer directoryWatcher.Close()

	go t.compileAndRun()

	directoryWatcher.Watch(dsnotify.DirectoryFunc(t.fileChangeEvent))
}

func (t *ServerTask) fileChangeEvent(e *fsnotify.Event, err error) {
	if err != nil {
		t.Say(t.writer, err.Error())
	}

	if t.currentCommand != nil {
		t.Say(t.writer, "") // newline
		t.kill()
	}

	t.compileAndRun()
}

func (t *ServerTask) kill() {
	defer func(begin time.Time) {
		t.Say(t.writer, fmt.Sprintf("time to kill process: %v", time.Since(begin)))
	}(time.Now())

	err := t.currentCommand.Process.Kill()
	if err != nil {
		t.Say(t.writer, err.Error())
	}
	t.currentCommand.Wait()
}

func (t *ServerTask) compile() {
	defer func(begin time.Time) {
		t.Say(t.writer, fmt.Sprintf("time to compile: %v", time.Since(begin)))
	}(time.Now())
	var (
		gitCommit, _ = exec.Command("git", "rev-parse", "--short", "HEAD").Output()
		buildNumber  = "1"
		version      = "v1.0.0"
		buildTime    = time.Now().UTC().Format(time.RFC3339Nano)
		cfgPkg       = "github.com/duhruh/blog/config"
	)

	cmd := exec.Command(
		"go",
		"build",
		"-o",
		t.outBinName,
		"-ldflags",
		""+t.ldflag(cfgPkg, "GitCommit", strings.Trim(string(gitCommit), "\n"))+" "+
			t.ldflag(cfgPkg, "BuildNumber", buildNumber)+" "+
			t.ldflag(cfgPkg, "Version", version)+" "+
			t.ldflag(cfgPkg, "BuildTime", buildTime),
		"cmd/api/main.go",
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

func (t *ServerTask) run() {
	t.Say(t.writer, "") // newline
	t.currentCommand = exec.Command(
		t.outBinName,
		"-http-bind-address=:8080",
		"-grpc-bind-address=:8081",
		"-environment=development",
	)
	t.currentCommand.Env = os.Environ()
	t.currentCommand.Stdin = os.Stdin
	t.currentCommand.Stdout = os.Stdout
	t.currentCommand.Stderr = os.Stderr
	err := t.currentCommand.Start()

	if err != nil {
		println(err.Error())
	}
}

func (t *ServerTask) compileAndRun() {
	t.compile()
	t.run()
}

func (t *ServerTask) ldflag(pkg string, variable string, value string) string {
	return "-X \"" + pkg + "." + variable + "=" + value + "\""
}
