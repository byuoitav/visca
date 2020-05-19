package visca

import (
	"context"
	"fmt"
)

const (
	_commandZoom = 0x07

	_funcZoomStop = 0x00
	_funcZoomTele = 0x02
	_funcZoomWide = 0x03
)

func (c *Camera) ZoomStop(ctx context.Context) error {
	p := make([]byte, 14)

	p[0] = 0x01
	p[1] = 0x00
	p[2] = 0x00
	p[3] = 0x06
	p[4] = 0x00
	p[5] = 0x00
	p[6] = 0x00
	p[7] = 0x03
	p[8] = 0x81
	p[9] = _command
	p[10] = _categoryCamera1
	p[11] = _commandZoom
	p[12] = _funcZoomStop
	p[13] = _terminator

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

	p := make([]byte, 14)

	p[0] = 0x01
	p[1] = 0x00
	p[2] = 0x00
	p[3] = 0x07
	p[4] = 0x00
	p[5] = 0x00
	p[6] = 0x00
	p[7] = _command
	p[8] = 0x81
	p[9] = _command
	p[10] = _categoryCamera1
	p[11] = _commandZoom
	p[12] = _funcZoomTele
	p[13] = _terminator

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

	p := make([]byte, 14)

	p[0] = 0x01
	p[1] = 0x00
	p[2] = 0x00
	p[3] = 0x07
	p[4] = 0x00
	p[5] = 0x00
	p[6] = 0x00
	p[7] = _command
	p[8] = 0x81
	p[9] = _command
	p[10] = _categoryCamera1
	p[11] = _commandZoom
	p[12] = _funcZoomWide
	p[13] = _terminator

	err = c.SendPayload(ctx, p)
	if err != nil {
		return err
	}

	return nil
}
