package visca

import (
	"encoding/binary"
	"errors"
)

type categoryCode byte

const (
	_categoryCodeInterface categoryCode = 0x00
	_categoryCodeCamera1   categoryCode = 0x04
	_categoryCodePanTilter categoryCode = 0x06
)

type command byte

const (
	_commandPanTiltDrive command = 0x01
	_commandMemory       command = 0x3f
	_commandZoom         command = 0x07
)

type payloadType [2]byte

var (
	_payloadTypeCommand = payloadType{0x01, 0x00}
	_payloadTypeInquiry = payloadType{0x01, 0x10}
)

const _terminator = 0xff

var ErrArgsTooLong = errors.New("payload arguments length must not exceed 0xffff")

type payload struct {
	Type payloadType

	IsInquiry    bool
	CategoryCode categoryCode
	Command      command
	Args         []byte
}

func (p payload) MarshalBinary() ([]byte, error) {
	length := uint16(len(p.Args) + 5)
	if length > 0xffff {
		return nil, ErrArgsTooLong
	}

	var buf []byte

	// bytes 0 and 1 are for the type
	// this function only works for typeCommand rn
	buf = append(buf, p.Type[0])
	buf = append(buf, p.Type[1])

	// bytes 2 and 3 are for the length
	lengthBuf := make([]byte, 2)
	binary.BigEndian.PutUint16(lengthBuf, length)
	buf = append(buf, lengthBuf...)

	// bytes 3-7 are for who knows what
	buf = append(buf, make([]byte, 4)...)

	// byte 8 is for the address
	buf = append(buf, 0x81)

	// byte 9 is either command (0x01) or inquiry (0x09)
	if p.IsInquiry {
		buf = append(buf, 0x09)
	} else {
		buf = append(buf, 0x01)
	}

	// byte 10 is the category code
	buf = append(buf, byte(p.CategoryCode))

	// byte 11 is the command
	buf = append(buf, byte(p.Command))

	// byte 12+ is the command arguments
	buf = append(buf, p.Args...)

	// last byte is the terminator
	buf = append(buf, _terminator)

	return buf, nil
}
