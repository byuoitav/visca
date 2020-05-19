package visca

import (
	"context"
	"fmt"
	"time"
)

const (
	_zoom     = 0x07
	_zoomStop = 0x00
	_zoomTele = 0x02
	_zoomWide = 0x03
)

func (c *Camera) ZoomStop(ctx context.Context) error {
	p := make([]byte, 16)

	p[0] = 0x01
	p[1] = 0x00
	p[2] = 0x00
	p[3] = 0x06
	p[4] = 0x00
	p[5] = 0x00
	p[6] = 0x00
	p[7] = 0x03
	p[8] = 0x81
	p[9] = 0x01
	p[10] = 0x04
	p[11] = _zoom
	p[12] = _zoomStop
	p[13] = 0xff

	err := c.SendPayload(ctx, p)
	if err != nil {
		return err
	}

	return nil
}

func (c *Camera) ZoomTele(ctx context.Context) error {
	err := c.ZoomStop(ctx)
	if err != nil {
		return fmt.Errorf("unable to stop zoom: %w", err)
	}

	p := make([]byte, 16)

	p[0] = 0x01
	p[1] = 0x00
	p[2] = 0x00
	p[3] = 0x07
	p[4] = 0x00
	p[5] = 0x00
	p[6] = 0x00
	p[7] = 0x01
	p[8] = 0x81
	p[9] = 0x01
	p[10] = 0x04
	p[11] = _zoom
	p[12] = _zoomTele
	p[13] = 0xff

	err = c.SendPayload(ctx, p)
	if err != nil {
		return err
	}

	return nil
}

func (c *Camera) ZoomWide(ctx context.Context) error {
	err := c.ZoomStop(ctx)
	if err != nil {
		return fmt.Errorf("unable to stop zoom: %w", err)
	}

	time.Sleep(100 * time.Millisecond)

	p := make([]byte, 16)

	p[0] = 0x01
	p[1] = 0x00
	p[2] = 0x00
	p[3] = 0x07
	p[4] = 0x00
	p[5] = 0x00
	p[6] = 0x00
	p[7] = 0x01
	p[8] = 0x81
	p[9] = 0x01
	p[10] = 0x04
	p[11] = _zoom
	p[12] = _zoomWide
	p[13] = 0xff

	err = c.SendPayload(ctx, p)
	if err != nil {
		return err
	}

	return nil
}
