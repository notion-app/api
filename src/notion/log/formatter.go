package log

import (
	"github.com/Sirupsen/logrus"
	"fmt"
	"strings"
)

const (
	LogLevelLength = "4"
	LogMessageLength = "60"
	LogKvLength = "10"
)

// Defines a logrus formatter which only outputs a message and fields
type NotionFormatter struct{}

func (nf NotionFormatter) Format(e *logrus.Entry) ([]byte, error) {
	line := fmt.Sprintf("%" + LogLevelLength + "." + LogLevelLength + "v | ", e.Level.String())
	line += fmt.Sprintf("%" + LogMessageLength + "." + LogMessageLength + "v | ", strings.Replace(e.Message, "\n", "", -1))
	for key, value := range e.Data {
		line += fmt.Sprintf("%" + LogKvLength + "." + LogKvLength + "v : ", strings.Replace(key, "\n", "", -1))
		line += fmt.Sprintf("%" + LogKvLength + "." + LogKvLength + "v | ", value)
	}
	line += "\n"
	return []byte(line), nil
}
