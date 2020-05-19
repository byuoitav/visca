package visca

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/byuoitav/connpool"
)

const (
	_command    = 0x01
	_terminator = 0xff
)

const (
	_categoryCamera1 = 0x04
	_categoryPanTilt = 0x06
)

type Camera struct {
	address string
	pool    *connpool.Pool

	logger Logger
	dialer net.Dialer
}

func New(addr string, opts ...Option) *Camera {
	options := options{
		ttl:    _defaultTTL,
		delay:  _defaultDelay,
		dialer: net.Dialer{},
	}

	for _, o := range opts {
		o.apply(&options)
	}

	cam := &Camera{
		address: addr,
		pool: &connpool.Pool{
			TTL:    options.ttl,
			Delay:  options.delay,
			Logger: options.logger,
		},
		logger: options.logger,
		dialer: options.dialer,
	}

	cam.pool.NewConnection = func(ctx context.Context) (net.Conn, error) {
		return cam.dialer.DialContext(ctx, "udp", cam.address)
	}

	return cam
}

func (c *Camera) SendPayload(ctx context.Context, p []byte) error {
	var resp []byte

	err := c.pool.Do(ctx, func(conn connpool.Conn) error {
		c.debugf("Sending payload: %# x", p)

		deadline, ok := ctx.Deadline()
		if !ok {
			deadline = time.Now().Add(3 * time.Second)
		}

		conn.SetWriteDeadline(deadline)

		n, err := conn.Write(p)
		switch {
		case err != nil:
			return fmt.Errorf("unable to write payload: %w", err)
		case n != len(p):
			return fmt.Errorf("unable to write payload: wrote %d/%d bytes", n, len(p))
		}

		resp, err = conn.ReadUntil(0xff, deadline)
		if err != nil {
			return fmt.Errorf("unable to read response: %w", err)
		}

		c.debugf("Got response: %# x", resp)
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
