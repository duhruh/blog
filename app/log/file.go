package log

import (
	"io"
	"encoding/json"
	"bytes"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
	"github.com/duhruh/blog/config"
)

func NewFileLogger(file string, l log.Logger, c config.ApplicationConfig) log.Logger{
	f, _ := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
	return fileLogger{file:f, next:l, config: c}
}



type fileLogger struct {
	file io.Writer
	next log.Logger
	config config.ApplicationConfig
}



func (fl fileLogger) Log(keyvals ...interface{}) error {
	var msg map[string]interface{}
	msg = make(map[string]interface{})


	for i := 0; i < len(keyvals); i += 2 {
		var (
			key = keyvals[i].(string)
			value = keyvals[i+1]
		)

		//if  key == "level" {
		//	continue
		//}

		var v interface{}
		switch value.(type) {
		case error:
			v = value.(error).Error()
			break
		case level.Value:
			v = value.(level.Value).String()
		default:
			v = value
		}
		msg[key] = v
	}

	var app map[string]interface{}
	app = make(map[string]interface{})

	app[fl.config.Name()] = msg

	js, _ := json.Marshal(app)

	var buf bytes.Buffer
	buf.Write(js)
	buf.WriteString("\n")

	fl.file.Write(buf.Bytes())

	return fl.next.Log(keyvals...)
}

