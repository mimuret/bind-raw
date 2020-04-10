package bind_raw

import (
	"github.com/miekg/dns"
	"golang.org/x/xerrors"
)

type Raw struct {
	header Header
	rrsets []RawRRSet
}

func (r *Raw) Unpack(bs []byte, off int) error {
	off, err := r.header.Unpack(bs, off)
	if err != nil {
		return xerrors.Errorf("failed to unpack header: %w", err)
	}
	for len(bs) > off {
		rrset := RawRRSet{}
		off, err = rrset.Unpack(bs, off)
		if err != nil {
			return xerrors.Errorf("failed to unpack rdata: %w", err)
		}
		r.rrsets = append(r.rrsets, rrset)
	}
	if len(bs) != off {
		return xerrors.Errorf("not finish read len=%d, off=%d", len(bs), off)
	}
	return nil
}

func (r *Raw) GetRRSets() []RawRRSet {
	return r.rrsets
}

func (r *Raw) GetRRs() ([]dns.RR, error) {
	var results []dns.RR
	for _, rrset := range r.rrsets {
		res, err := rrset.GetRRs()
		if err != nil {
			return nil, err
		}
		results = append(results, res...)
	}
	return results, nil
}

func (r *Raw) GetRFC3597s() []*dns.RFC3597 {
	var results []*dns.RFC3597
	for _, rrset := range r.rrsets {
		res := rrset.GetRFC3597s()
		results = append(results, res...)
	}
	return results
}
