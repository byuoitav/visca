package visca

import (
	"context"
)

const (
	_funcMemoryReset  = 0x00
	_funcMemorySet    = 0x01
	_funcMemoryRecall = 0x02

	_memoryChannelMin = 0x00
	_memoryChannelMax = 0x7f
)

func (c *Camera) MemoryRecall(ctx context.Context, channel byte) error {
	payload := payload{
		Type:         _payloadTypeCommand,
		IsInquiry:    false,
		CategoryCode: _categoryCodeCamera1,
		Command:      _commandMemory,
		Args: []byte{
			_funcMemoryRecall,
			channel,
		},
	}

	if err := c.sendPayload(ctx, payload); err != nil {
		return err
	}

	return nil
}
