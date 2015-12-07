package log

import (
	"github.com/Sirupsen/logrus"
	"fmt"
)

const (
	LogLevelLength = "4"
	LogMessageLength = "45"
	LogKvLength = "10"
)

// Defines a logrus formatter which only outputs a message and fields
type NotionFormatter struct{}

func (nf NotionFormatter) Format(e *logrus.Entry) ([]byte, error) {
	line := fmt.Sprintf("%" + LogLevelLength + "." + LogLevelLength + "v | ", e.Level.String())
	line += fmt.Sprintf("%" + LogMessageLength + "." + LogMessageLength + "v | ", e.Message)
	for key, value := range e.Data {
		line += fmt.Sprintf("%" + LogKvLength + "." + LogKvLength + "v : ", key)
		line += fmt.Sprintf("%" + LogKvLength + "." + LogKvLength + "v | ", value)
	}
	line += "\n"
	return []byte(line), nil
}
