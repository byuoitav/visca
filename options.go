package visca

import (
	"net"
	"time"
)

var (
	_defaultTTL   = 30 * time.Second
	_defaultDelay = 75 * time.Millisecond
)

type options struct {
	ttl    time.Duration
	delay  time.Duration
	logger Logger
	dialer net.Dialer
}

// Option configures how we create the DSP.
type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

// WithTTL changes the TTL for the underlying UDP connection to the DSP.
// The default value is 30 seconds.
// See more details about TTL in https://github.com/byuoitav/connpool.
func WithTTL(t time.Duration) Option {
	return optionFunc(func(o *options) {
		o.ttl = t
	})
}

// WithDelay changes the delay between sending commands to the DSP.
// The default value is 250 milliseconds.
// See more details about delay in https://github.com/byuoitav/connpool.
func WithDelay(t time.Duration) Option {
	return optionFunc(func(o *options) {
		o.delay = t
	})
}

// WithLogger adds a logger to DSP.
// DSP will log appropriate information about the underlying connection and the commands being sent.
// The default value is nil, meaning that no logs are written.
func WithLogger(l Logger) Option {
	return optionFunc(func(o *options) {
		o.logger = l
	})
}

func WithDialer(d net.Dialer) Option {
	return optionFunc(func(o *options) {
		o.dialer = d
	})
}
