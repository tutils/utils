package addr

import (
	"net"
	"testing"
)

func TestRealAddress(t *testing.T) {
	addrStr := ":https"
	addr, err := net.ResolveTCPAddr("tcp", addrStr)
	if err != nil {
		t.Fatal(err)
	}

	res, err := RealAddress(addr)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s -> %s", addrStr, res)
}
