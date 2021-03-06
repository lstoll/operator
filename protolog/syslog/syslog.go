// Package protologsyslog defines functionality for integration with syslog.
package protologsyslog

import (
	"log/syslog"

	"github.com/sr/operator/protolog"
)

var (
	// DefaultTextMarshaller is the default text Marshaller for syslog.
	DefaultTextMarshaller = protolog.NewTextMarshaller(
		protolog.TextMarshallerDisableTime(),
	)
)

// PusherOption is an option for constructing a new Pusher.
type PusherOption func(*pusher)

// PusherWithMarshaller uses the Marshaller for the Pusher.
//
// By default, DefaultTextMarshaller is used.
func PusherWithMarshaller(marshaller protolog.Marshaller) PusherOption {
	return func(pusher *pusher) {
		pusher.marshaller = marshaller
	}
}

// NewPusher creates a new protolog.Pusher that logs using syslog.
func NewPusher(writer *syslog.Writer, options ...PusherOption) protolog.Pusher {
	return newPusher(writer, options...)
}
