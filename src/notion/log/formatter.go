package log

import (
	"github.com/Sirupsen/logrus"
)

const (
	LogLineLength = 45
)

// Defines a logrus formatter which only outputs a message and fields
type NotionFormatter struct{}

func (nf NotionFormatter) Format(e *logrus.Entry) ([]byte, error) {
	line := ""
	if len(e.Message) >= LogLineLength {
		line += e.Message[:LogLineLength] + " | "
	} else {
		line += e.Message
		for i := 0; i < (LogLineLength - len(e.Message)); i += 1 {
			line += " "
		}
		line += " | "
	}
	line += "\n"
	return []byte(line), nil
}
