package bind_raw

import (
	"io"
	"io/ioutil"
)

func Unpack(r io.Reader) (*Raw, error) {
	var err error
	off := 0
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	raw := Raw{}
	if err := raw.Unpack(bs, off); err != nil {
		return nil, err
	}
	return &raw, nil
}

func unpackUint16(msg []byte, off int) (uint16, int, error) {
	if off+uint16Len > len(msg) {
		return 0, off, errBaseLen
	}
	return uint16(msg[off])<<8 | uint16(msg[off+1]), off + uint16Len, nil
}

func skipUint16(msg []byte, off int) (int, error) {
	if off+uint16Len > len(msg) {
		return off, errBaseLen
	}
	return off + uint16Len, nil
}

func unpackUint32(msg []byte, off int) (uint32, int, error) {
	if off+uint32Len > len(msg) {
		return 0, off, errBaseLen
	}
	v := uint32(msg[off])<<24 | uint32(msg[off+1])<<16 | uint32(msg[off+2])<<8 | uint32(msg[off+3])
	return v, off + uint32Len, nil
}

func skipUint32(msg []byte, off int) (int, error) {
	if off+uint32Len > len(msg) {
		return off, errBaseLen
	}
	return off + uint32Len, nil
}
