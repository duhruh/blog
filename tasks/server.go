package tasks

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"time"

	"github.com/duhruh/tackle/dsnotify"
	"github.com/duhruh/tackle/task"
	"github.com/fsnotify/fsnotify"
	"strconv"
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
	buildNumber    int
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

	t.outBinName = fmt.Sprintf("dist/%s_blog", time.Now().UTC().Format("20060102"))

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

	t.buildNumber++
	cmd := exec.Command(
		"go",
		"run",
		"cmd/task/main.go",
		"build",
		"--version=v0.0.0-alpha."+strconv.Itoa(t.buildNumber),
		"--build="+strconv.Itoa(t.buildNumber),
		"dist/",
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
		"-config=config/app.yml",
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
