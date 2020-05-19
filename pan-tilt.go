package visca

import (
	"context"
	"errors"
)

const (
	_commandPanTiltDrive = 0x01

	TiltDirectionUp   = 0x01
	TiltDirectionDown = 0x02
	TiltDirectionStop = 0x03

	PanDirectionLeft  = 0x01
	PanDirectionRight = 0x02
	PanDirectionStop  = 0x03

	PanTiltSpeedMin = 0x00
	PanTiltSpeedMax = 0x18
)

var (
	ErrInvalidPanTiltSpeed     = errors.New("speed out of range")
	ErrInvalidPanTiltDirection = errors.New("invalid pan/tilt direction")
)

// TODO check if it's out of range?

func (c *Camera) TiltUp(ctx context.Context, speed byte) error {
	return c.PanTiltDrive(ctx, PanDirectionStop, TiltDirectionUp, PanTiltSpeedMin, speed)
}

func (c *Camera) TiltDown(ctx context.Context, speed byte) error {
	return c.PanTiltDrive(ctx, PanDirectionStop, TiltDirectionDown, PanTiltSpeedMin, speed)
}

func (c *Camera) PanLeft(ctx context.Context, speed byte) error {
	return c.PanTiltDrive(ctx, PanDirectionLeft, TiltDirectionStop, speed, PanTiltSpeedMin)
}

func (c *Camera) PanRight(ctx context.Context, speed byte) error {
	return c.PanTiltDrive(ctx, PanDirectionRight, TiltDirectionStop, speed, PanTiltSpeedMin)
}

func (c *Camera) PanTiltStop(ctx context.Context) error {
	return c.PanTiltDrive(ctx, PanDirectionStop, TiltDirectionStop, PanTiltSpeedMin, PanTiltSpeedMin)
}

func (c *Camera) PanTiltDrive(ctx context.Context, panDir, tiltDir, panSpeed, tiltSpeed byte) error {
	switch {
	case panSpeed < PanTiltSpeedMin || panSpeed > PanTiltSpeedMax:
		return ErrInvalidPanTiltSpeed
	case tiltSpeed < PanTiltSpeedMin || tiltSpeed > PanTiltSpeedMax:
		return ErrInvalidPanTiltSpeed
	case panDir != PanDirectionLeft && panDir != PanDirectionRight && panDir != PanDirectionStop:
		return ErrInvalidPanTiltDirection
	case tiltDir != TiltDirectionUp && tiltDir != TiltDirectionDown && tiltDir != TiltDirectionStop:
		return ErrInvalidPanTiltDirection
	}

	p := make([]byte, 17)

	p[0] = 0x01
	p[1] = 0x00
	p[2] = 0x00
	p[3] = 0x09
	p[4] = 0x00
	p[5] = 0x00
	p[6] = 0x00
	p[7] = 0x01
	p[8] = 0x81
	p[9] = _command
	p[10] = _categoryPanTilter
	p[11] = _commandPanTiltDrive
	p[12] = panSpeed
	p[13] = tiltSpeed
	p[14] = panDir
	p[15] = tiltDir
	p[16] = _terminator

	err := c.SendPayload(ctx, p)
	if err != nil {
		return err
	}

	return nil
}
