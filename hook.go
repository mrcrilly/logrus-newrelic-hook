package logrusnewrelic

import (
	"github.com/sirupsen/logrus"
)

var DefaultLevels = []logrus.Level{
	logrus.FatalLevel,
	logrus.ErrorLevel,
	logrus.WarnLevel,
	logrus.InfoLevel,
	logrus.DebugLevel,
}

type Hook struct {
	Client *Client
	levels []logrus.Level
}

func NewHook(client *Client, levels []logrus.Level) *Hook {
	return &Hook{client, levels}
}

func (h *Hook) Fire(entry *logrus.Entry) error {
	_, err := h.Client.Log(entry)
	return err
}

func (h *Hook) Levels() []logrus.Level {
	if h.levels == nil {
		return DefaultLevels
	}
	return h.levels
}
