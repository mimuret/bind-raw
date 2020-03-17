package bind_raw

import "golang.org/x/xerrors"
import 	"github.com/miekg/dns"

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
		r.rrsets = append(r.rrsets,rrset)
	}
	if len(bs) != off {
		return xerrors.Errorf("not finish read len=%d, off=%d", len(bs),off)
	}
	return nil
}
func (r *Raw)GetRRs() []dns.RR {
	var results []dns.RR
	for _,rrset := range r.rrsets {
		res := rrset.GetRRs()
		results = append(results,res...)
	}
	return results
}
