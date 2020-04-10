package bind_raw

import (
	"encoding/hex"
	"errors"

	"github.com/miekg/dns"
	"golang.org/x/xerrors"
)

var (
	uint16Len  = 2
	uint32Len  = 4
	errBaseLen = errors.New("insufficient data for base length type")
)

type Rdataset struct {
	TotalLen uint32
	Class    uint16
	Type     uint16
	Covers   uint16
	TTL      uint32
	NRData   uint32
}

func (r *Rdataset) Unpack(bs []byte, off int) (int, error) {
	var err error
	r.TotalLen, off, err = unpackUint32(bs, off)
	if err != nil {
		return off, xerrors.Errorf("failed to parse total len: %w", err)
	}
	r.Class, off, err = unpackUint16(bs, off)
	if err != nil {
		return off, xerrors.Errorf("failed to parse class: %w", err)
	}
	r.Type, off, err = unpackUint16(bs, off)
	if err != nil {
		return off, xerrors.Errorf("failed to parse type: %w", err)
	}
	r.Covers, off, err = unpackUint16(bs, off)
	if err != nil {
		return off, xerrors.Errorf("failed to parse covers: %w", err)
	}
	r.TTL, off, err = unpackUint32(bs, off)
	if err != nil {
		return off, xerrors.Errorf("failed to parse ttl: %w", err)
	}
	r.NRData, off, err = unpackUint32(bs, off)
	if err != nil {
		return off, xerrors.Errorf("failed to parse rdata count: %w", err)
	}
	return off, nil
}

// dump_rdatasets_raw
// dump_rdataset_raw
type RawRRSet struct {
	dataset Rdataset
	name    string
	Rdatas  [][]byte
}

func (r *RawRRSet) Unpack(bs []byte, off int) (int, error) {
	var err error
	off, err = r.dataset.Unpack(bs, off)
	if err != nil {
		return off, xerrors.Errorf("failt to parse rdataset: %w", err)
	}
	_, off, err = unpackUint16(bs, off)
	r.name, off, err = dns.UnpackDomainName(bs, off)
	if err != nil {
		return off, xerrors.Errorf("failt to parse name: %w", err)
	}

	var c uint32
	for ; c < r.dataset.NRData; c++ {
		var rdlength uint16
		rdlength, off, err = unpackUint16(bs, off)
		if err != nil {
			return off, xerrors.Errorf("failed to parse rdlength: %w", err)
		}
		r.Rdatas = append(r.Rdatas, bs[off:off+int(rdlength)])
		off += int(rdlength)
	}

	return off, nil
}

func (r *RawRRSet) GetRRs() ([]dns.RR, error) {
	var results []dns.RR
	for _, rrdata := range r.Rdatas {
		hexBytes := make([]byte, len(rrdata)*2)
		hex.Encode(hexBytes, rrdata)
		header := dns.RR_Header{r.name, r.dataset.Type, r.dataset.Class, r.dataset.TTL, uint16(len(rrdata))}
		rr, _, err := dns.UnpackRRWithHeader(header, rrdata, 0)
		if err != nil {
			return nil, err
		}
		results = append(results, rr)
	}
	return results, nil
}

func (r *RawRRSet) GetRFC3597s() []*dns.RFC3597 {
	var results []*dns.RFC3597
	for _, rrdata := range r.Rdatas {
		hexBytes := make([]byte, len(rrdata)*2)
		hex.Encode(hexBytes, rrdata)
		rr := dns.RFC3597{
			Hdr:   dns.RR_Header{r.name, r.dataset.Type, r.dataset.Class, r.dataset.TTL, uint16(len(rrdata))},
			Rdata: string(hexBytes),
		}
		results = append(results, &rr)
	}
	return results
}
