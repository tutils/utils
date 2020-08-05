package basen

import "testing"

func TestConverter_ToBaseN(t *testing.T) {
	c := NewConverter(Base62CharSet)
	t.Log(string(c.ToBaseN(4354536116143)))
}

func TestConverter_ToNumber(t *testing.T) {
	c := NewConverter(Base62CharSet)
	n, err := c.ToNumber([]byte("BOpKlv4L"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(n)

	_, err = c.ToNumber([]byte("BOpKlv4L="))
	if err == nil {
		t.Fail()
	}
}
