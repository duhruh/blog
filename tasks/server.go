package tasks

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/duhruh/tackle"
	"github.com/duhruh/tackle/dsnotify"
	"github.com/duhruh/tackle/task"
	"github.com/fsnotify/fsnotify"

	"github.com/duhruh/blog/app"
	"github.com/duhruh/blog/config"
)

const (
	DefaultDistDir = "dist"
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
	config         config.ApplicationConfig
	env            tackle.Environment
}

func NewServerTask() task.Task {
	return &ServerTask{
		Helpers:          task.NewHelpers(),
		name:             "server",
		shortDescription: "Runs the tackle server",
		description:      "Runs the tackle server",
		options: []task.Option{
			task.NewOption("watch", "recompiles on file change"),
			task.NewOption("http", "if set will start a http server listening on the specified port"),
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
	t.env = tackle.Development
	t.config = app.NewConfigFromYamlFile(t.env, configPath)

	t.Say(w, "\nWelcome to Tackle v1.0.0\n")

	option, _ := t.GetOption(t.options, "watch")
	if option.Value() == nil {
		t.compileAndRun()
		t.currentCommand.Wait()
	}

	htt, _ := t.GetOption(t.options, "http")
	if htt.Value() != nil {
		go func() {
			http.HandleFunc("/recompile", t.httpRecompile)
			http.ListenAndServe(htt.Value().(string), nil)
		}()
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

type ByFileDateName []string

func (s ByFileDateName) Len() int {
	return len(s)
}
func (s ByFileDateName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByFileDateName) Less(i, j int) bool {

	aFile := s.getDateFromFileName(s[i])
	bFile := s.getDateFromFileName(s[j])

	return aFile.After(bFile)
}

func (s ByFileDateName) getDateFromFileName(name string) time.Time {
	parts := strings.Split(name, "_")

	t, err := time.Parse(BuildTimeDateFormat, parts[0])
	if err != nil {
		panic(err)
	}

	return t
}

func (t *ServerTask) getBinName() string {
	files, err := ioutil.ReadDir(DefaultDistDir)
	if err != nil {
		panic(err)
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	sort.Sort(ByFileDateName(fileNames))

	return fileNames[0]
}

func (t *ServerTask) fileChangeEvent(e *fsnotify.Event, err error) {
	if err != nil {
		t.Say(t.writer, err.Error())
	}

	t.recompile()
}

func (t *ServerTask) httpRecompile(res http.ResponseWriter, req *http.Request) {
	t.recompile()
	res.WriteHeader(200)
}

func (t *ServerTask) recompile() {
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
		t.config.TaskEntryPoint(),
		"build",
		fmt.Sprintf("--build=%s", strconv.Itoa(t.buildNumber)),
		DefaultDistDir,
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
		filepath.Join(DefaultDistDir, t.getBinName()),
		fmt.Sprintf("-config=%s", configPath),
		fmt.Sprintf("-environment=%v", t.env),
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
