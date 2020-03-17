package bind_raw

import (
	"os"
	"testing"
	"fmt"
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
	fmt.Println(raw)
	rrs := raw.GetRRs()  
	for _,rr :=  range rrs {
		fmt.Println(rr)
	}
}
