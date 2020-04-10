package bind_raw

import (
	"fmt"
	"os"
	"testing"
)

func TestUnpack(t *testing.T) {
	r, err := os.Open("./test/example.raw")
	if err != nil {
		t.Fatal(err)
	}
	raw, err := Unpack(r)
	if err != nil {
		t.Fatal(err)
	}
	rrs := raw.GetRFC3597s()
	for _, rr := range rrs {
		fmt.Println(rr)
	}
}
