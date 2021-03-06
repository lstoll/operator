package protologsyslog

import (
	"log/syslog"

	"github.com/sr/operator/protolog"
)

var (
	levelToLogFunc = map[protolog.Level]func(*syslog.Writer, string) error{
		protolog.LevelDebug: (*syslog.Writer).Debug,
		protolog.LevelInfo:  (*syslog.Writer).Info,
		protolog.LevelError: (*syslog.Writer).Err,
	}
)

type pusher struct {
	writer     *syslog.Writer
	marshaller protolog.Marshaller
}

func newPusher(writer *syslog.Writer, options ...PusherOption) *pusher {
	pusher := &pusher{writer, DefaultTextMarshaller}
	for _, option := range options {
		option(pusher)
	}
	return pusher
}

func (p *pusher) Flush() error {
	return nil
}

func (p *pusher) Push(entry *protolog.Entry) error {
	data, err := p.marshaller.Marshal(entry)
	if err != nil {
		return err
	}
	return levelToLogFunc[entry.Level](p.writer, string(data))
}
