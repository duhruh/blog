package log

import (
	"io"

	"strings"

	"bytes"
	"fmt"
	"github.com/fatih/color"
	//"regexp"
	"regexp"
)

type colorWriter struct {
	io.Writer
}

type colorSprinter func(a ...interface{}) string

var logLevelColorMap = map[string]colorSprinter{
	"info":    color.New(color.FgCyan).SprintFunc(),
	"warn":    color.New(color.FgYellow).SprintFunc(),
	"error":   color.New(color.FgRed).SprintFunc(),
	"debug":   color.New(color.FgHiBlue).SprintFunc(),
	"default": color.New(color.FgGreen).SprintFunc(),
}

func NewColorWriter(writer io.Writer) io.Writer {
	return colorWriter{Writer: writer}
}

func (cw colorWriter) Write(p []byte) (n int, err error) {
	var (
		parts     []string
		level     string
		timestamp string
		message   string
		prefix    string
		keyValues bytes.Buffer
		out       bytes.Buffer
		colorFunc colorSprinter
	)

	parts = cw.ripString(string(p))

	colorFunc, level, parts = cw.pluckLevel(parts)
	message, parts = cw.pluckMessage(parts)
	timestamp, parts = cw.pluckTimestamp(parts)

	keyValues = cw.formatKeyValues(parts, colorFunc)

	prefix = cw.formatLogMessage(level, timestamp, message, colorFunc)

	out.WriteString(prefix + "\t" + keyValues.String() + "\n")

	return cw.Writer.Write(out.Bytes())
}

// rip string converts the raw log message to its key value pairs
func (cw colorWriter) ripString(str string) []string {
	parts := strings.Split(str, " ")
	for i, k := range parts {
		if !strings.Contains(k, "=") {
			parts[i-1] = parts[i-1] + " " + k
			parts = append(parts[:i], parts[i+1:]...) //remove it
		}
	}

	return parts
}

// removes the given index from the given array
func (cw colorWriter) removeFromArray(arr []string, index int) []string {
	end := index + 1

	if end > len(arr) {
		return arr[:index]
	}
	return append(arr[:index], arr[end:]...)
}

// find the value in the key value pair array
func (cw colorWriter) pluck(k string, arr []string) (string, int) {
	for i, p := range arr {
		key, val := cw.stringToKeyValue(p)

		if key == k {
			return val, i
		}
	}
	return "", 0
}

func (cw colorWriter) pluckMessage(arr []string) (string, []string) {
	message, i := cw.pluck("message", arr)

	var re = regexp.MustCompile(`(?:^")?(.*?)(?:")?$`)
	matches := re.FindStringSubmatch(message)
	message = matches[0]

	if len(matches) > 0 {
		message = matches[1]
	}

	if message == "" {
		return message, arr
	}

	return message, cw.removeFromArray(arr, i)
}

func (cw colorWriter) pluckTimestamp(arr []string) (string, []string) {
	timestamp, i := cw.pluck("timestamp", arr)
	if timestamp == "" {
		return "0000", arr
	}

	return timestamp, cw.removeFromArray(arr, i)
}
func (cw colorWriter) pluckLevel(arr []string) (colorSprinter, string, []string) {
	levelName, i := cw.pluck("level", arr)

	if levelName == "" {
		levelName = "default"
	}
	colorFunc := logLevelColorMap[levelName]
	level := strings.ToUpper(levelName)[0:4]

	return colorFunc, level, cw.removeFromArray(arr, i)
}

func (cw colorWriter) formatKeyValues(arr []string, col colorSprinter) bytes.Buffer {
	var keyValues bytes.Buffer
	for _, k := range arr {
		key, val := cw.stringToKeyValue(k)

		keyValues.WriteString(fmt.Sprintf("%s=%s ", col(key), val))
	}

	return keyValues
}

func (cw colorWriter) stringToKeyValue(str string) (string, string) {
	keyval := strings.SplitN(str, "=", 2)

	return strings.TrimSpace(keyval[0]), strings.TrimSpace(keyval[1])
}

func (cw colorWriter) formatLogMessage(level string, timestamp string, message string, col colorSprinter) string {
	return fmt.Sprintf("%s[%s] %s", col(level), timestamp, message)
}
