
// This code is taken from github.com/weekface/mgorus and modified because
// I dont like the way he modifies the log entry inside a hook. I also removed
// the need to specify a database because it should be already included in the mgoUrl
// passed in.
package log

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoHook struct {
	c *mgo.Collection
}
type M bson.M

func NewMongoHook(mgoUrl, collection string) (*MongoHook, error) {
	session, err := mgo.Dial(mgoUrl)
	if err != nil {
		return nil, err
	}
	return &MongoHook{c: session.DB("").C(collection)}, nil
}

func (h *MongoHook) Fire(entry *logrus.Entry) error {
	entryData := map[string]interface{}{
		"level": entry.Level.String(),
		"time": entry.Time,
		"message": entry.Message,
	}
	for k, v := range entry.Data {
		entryData[k] = v
	}
	mgoErr := h.c.Insert(M(entryData))
	if mgoErr != nil {
		return fmt.Errorf("Failed to send log entry to mongodb: %s", mgoErr)
	}
	return nil
}

func (h *MongoHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}
