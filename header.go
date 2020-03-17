package bind_raw

import (
	"golang.org/x/xerrors"
)

type Header struct {
	Format       uint32
	Version      uint32
	Dumptime     uint32
	Flags        uint32
	SourceSerial uint32
	LastXfrIn    uint32
}

func (h *Header) Unpack(bs []byte, off int) (int, error) {
	var err error
	h.Format, off, err = unpackUint32(bs, off)
	if err != nil {
		return off, xerrors.Errorf("failed to parse header format: %w", err)
	}
	h.Version, off, err = unpackUint32(bs, off)
	if err != nil {
		return off, xerrors.Errorf("failed to parse header version: %w", err)
	}
	h.Dumptime, off, err = unpackUint32(bs, off)
	if err != nil {
		return off, xerrors.Errorf("failed to parse header dumptime: %w", err)
	}
	h.Flags, off, err = unpackUint32(bs, off)
	if err != nil {
		return off, xerrors.Errorf("failed to parse header flags: %w", err)
	}
	h.SourceSerial, off, err = unpackUint32(bs, off)
	if err != nil {
		return off, xerrors.Errorf("failed to parse header sourceserial: %w", err)
	}
	h.LastXfrIn, off, err = unpackUint32(bs, off)
	if err != nil {
		return off, xerrors.Errorf("failed to parse header lastxfrin: %w", err)
	}
	return off, nil
}

