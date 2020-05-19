package visca

import "context"

const (
	_commandMemory = 0x3f

	_funcMemoryReset  = 0x00
	_funcMemorySet    = 0x01
	_funcMemoryRecall = 0x02

	_memoryChannelMin = 0x00
	_memoryChannelMax = 0x7f
)

func (c *Camera) MemoryRecall(ctx context.Context, channel byte) error {
	p := make([]byte, 15)

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
	p[11] = _commandMemory
	p[12] = _funcMemoryRecall
	p[13] = channel
	p[14] = _terminator

	err := c.SendPayload(ctx, p)
	if err != nil {
		return err
	}

	return nil
}
