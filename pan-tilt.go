package visca

import (
	"context"
)

const (
	_panTilt = 0x01

	_tiltUp       = 0x01
	_tiltDown     = 0x02
	_tiltStop     = 0x03
	_tiltSpeedMin = 0x01
	_tiltSpeedMax = 0x14

	_panLeft     = 0x01
	_panRight    = 0x02
	_panStop     = 0x03
	_panSpeedMin = 0x01
	_panSpeedMax = 0x18
)

func (c *Camera) TiltUp(ctx context.Context, speed byte) error {
	p := make([]byte, 17)

	p[0] = 0x01
	p[1] = 0x00
	p[2] = 0x00
	p[3] = 0x09
	p[4] = 0x00
	p[5] = 0x00
	p[6] = 0x00
	p[7] = _command
	p[8] = 0x81
	p[9] = 0x01
	p[10] = _categoryPanTilt
	p[11] = _panTilt
	p[12] = _panSpeedMin
	p[13] = speed
	p[14] = _panStop
	p[15] = _tiltUp
	p[16] = _terminator

	err := c.SendPayload(ctx, p)
	if err != nil {
		return err
	}

	return nil
}

func (c *Camera) TiltDown(ctx context.Context, speed byte) error {
	p := make([]byte, 17)

	p[0] = 0x01
	p[1] = 0x00
	p[2] = 0x00
	p[3] = 0x09
	p[4] = 0x00
	p[5] = 0x00
	p[6] = 0x00
	p[7] = _command
	p[8] = 0x81
	p[9] = 0x01
	p[10] = _categoryPanTilt
	p[11] = _panTilt
	p[12] = _panSpeedMin
	p[13] = speed
	p[14] = _panStop
	p[15] = _tiltDown
	p[16] = _terminator

	err := c.SendPayload(ctx, p)
	if err != nil {
		return err
	}

	return nil
}

func (c *Camera) PanLeft(ctx context.Context, speed byte) error {
	p := make([]byte, 17)

	p[0] = 0x01
	p[1] = 0x00
	p[2] = 0x00
	p[3] = 0x09
	p[4] = 0x00
	p[5] = 0x00
	p[6] = 0x00
	p[7] = _command
	p[8] = 0x81
	p[9] = 0x01
	p[10] = _categoryPanTilt
	p[11] = _panTilt
	p[12] = speed
	p[13] = _tiltSpeedMin
	p[14] = _panLeft
	p[15] = _tiltStop
	p[16] = _terminator

	err := c.SendPayload(ctx, p)
	if err != nil {
		return err
	}

	return nil
}

func (c *Camera) PanRight(ctx context.Context, speed byte) error {
	p := make([]byte, 17)

	p[0] = 0x01
	p[1] = 0x00
	p[2] = 0x00
	p[3] = 0x09
	p[4] = 0x00
	p[5] = 0x00
	p[6] = 0x00
	p[7] = _command
	p[8] = 0x81
	p[9] = 0x01
	p[10] = _categoryPanTilt
	p[11] = _panTilt
	p[12] = speed
	p[13] = _tiltSpeedMin
	p[14] = _panRight
	p[15] = _tiltStop
	p[16] = _terminator

	err := c.SendPayload(ctx, p)
	if err != nil {
		return err
	}

	return nil
}

func (c *Camera) PanTiltStop(ctx context.Context) error {
	p := make([]byte, 17)

	p[0] = 0x01
	p[1] = 0x00
	p[2] = 0x00
	p[3] = 0x09
	p[4] = 0x00
	p[5] = 0x00
	p[6] = 0x00
	p[7] = _command
	p[8] = 0x81
	p[9] = 0x01
	p[10] = _categoryPanTilt
	p[11] = _panTilt
	p[12] = _panSpeedMin
	p[13] = _tiltSpeedMin
	p[14] = _panStop
	p[15] = _tiltStop
	p[16] = _terminator

	err := c.SendPayload(ctx, p)
	if err != nil {
		return err
	}

	return nil
}
