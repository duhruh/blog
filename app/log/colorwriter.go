package log

import (
	"io"

	"strings"

	"github.com/fatih/color"
)

type colorWriter struct {
	io.Writer
}

var logLevelColorMap = map[string]*color.Color{
	"info":    color.New(color.FgCyan),
	"warn":    color.New(color.FgYellow),
	"default": color.New(color.FgGreen),
}

func NewColorWriter(writer io.Writer) io.Writer {
	return colorWriter{Writer: writer}
}

func (cw colorWriter) Write(p []byte) (n int, err error) {
	str := string(p)

	parts := cw.ripString(str)
	//col := logLevelColorMap["default"]
	//levelName := "INFO"
	//message := ""
	for i, k := range parts {
		keyval := strings.SplitN(k, "=", 2)
		if keyval[0] == "level" {
			//col = logLevelColorMap[keyval[1]]
			//levelName = strings.ToUpper(keyval[1])
			parts = append(parts[:i], parts[i+1:]...) // remove it
			continue
		}
		if keyval[0] == "message" {
			//message = keyval[1]
			//parts = append(parts[:i], parts[i+1:]...) // remove it
			continue
		}

	}
	//red := color.New(color.FgRed).SprintFunc()
	//uhh := fmt.Sprintf("%s[%s] %s ", red("RED"))

	//cw.Writer.Write([]byte(uhh))

	//return col.Fprint(cw.Writer, str)

	return cw.Writer.Write(p)
}

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
