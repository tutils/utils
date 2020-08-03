package addr

import (
	"testing"
)

func TestRealAddress(t *testing.T) {
	addr := ":8080"
	res, err := RealAddress(addr)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s -> %s", addr, res)
}
