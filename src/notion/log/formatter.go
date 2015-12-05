
package log

import (
  "github.com/Sirupsen/logrus"
)

// Defines a logrus formatter which only outputs a message and fields
type NotionFormatter struct {}

func (nf NotionFormatter) Format(e *logrus.Entry) ([]byte, error) {
  return nil, nil
}
