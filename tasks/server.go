package tasks

import (
	"io"

	"github.com/duhruh/tackle/dsnotify"
	"github.com/duhruh/tackle/task"
	"github.com/fsnotify/fsnotify"
	"os"
	"os/exec"

	"fmt"
	"regexp"
	"strings"
	"time"
)

type ServerTask struct {
	task.Helpers
	shortDescription string
	description      string
	name             string
	options          []task.Option
	arguments        []task.Argument

	currentCommand *exec.Cmd
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

	option, _ := t.GetOption(t.options, "watch")
	if option.Value() == nil {
		t.build()
		t.currentCommand.Wait()
	}

	ww, _ := dsnotify.NewDirectoryWatcher()

	ww.IgnoreRegex(regexp.MustCompile(`^.glide`))
	ww.IgnoreRegex(regexp.MustCompile(`^.git`))
	ww.AddDirectory("./")
	ww.RegisterFileRegex(regexp.MustCompile(`(.+\.go)`))

	defer ww.FsWatcher().Close()

	go t.build()

	ww.Watch(dsnotify.DirectoryFunc(func(e *fsnotify.Event, err error) {

		if err != nil {
			t.Say(w, err.Error())
		}

		if t.currentCommand != nil {
			err := t.currentCommand.Process.Kill()
			if err != nil {
				t.Say(w, err.Error())
			}
			t.currentCommand.Wait()

		}

		t.Say(w, "rebuild...")
		t.build()
	}))
}

func (t *ServerTask) build() {
	var (
		gitCommit, _ = exec.Command("git", "rev-parse", "--short", "HEAD").Output()
		buildNumber  = "1"
		version      = "v1.0.0"
		buildTime    = time.Now().UTC().Format(time.RFC3339Nano)
		cfgPkg       = "github.com/duhruh/blog/config"
		bin          = fmt.Sprintf("/tmp/go-bin-%s", time.Now().String())
	)

	cmd := exec.Command(
		"go",
		"build",
		"-o",
		bin,
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

	t.currentCommand = exec.Command(
		bin,
		"-http-bind-address=:8080",
		"-grpc-bind-address=:8081",
	)
	t.currentCommand.Env = os.Environ()
	t.currentCommand.Stdin = os.Stdin
	t.currentCommand.Stdout = os.Stdout
	t.currentCommand.Stderr = os.Stderr
	err = t.currentCommand.Start()

	if err != nil {
		println(err.Error())
	}
}

func (t *ServerTask) ldflag(pkg string, variable string, value string) string {
	return "-X \"" + pkg + "." + variable + "=" + value + "\""
}
