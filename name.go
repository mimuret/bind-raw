package bind_raw

type IscName struct {
	Magic      uint
	Ndata      []byte
	Length     uint
	Labels     uint
	Attributes uint
	Offsets    []byte
}

func (r *IscName) Unpack(bs []byte, off int) (int, error) {
	return off, nil
}
