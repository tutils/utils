package snowflake

import (
	"testing"
	"time"
)

func float64Safe(id ID) bool {
	return int64(float64(id)) == int64(id)
}

func TestIDGen_Generate(t *testing.T) {
	gen, err := NewIDGen(WithNode(0, 3), WithStepBits(9))
	if err != nil {
		t.Fatal(err)
	}

	const maxNumber = 100
	var idList [maxNumber]ID
	for i := 0; i < maxNumber; i++ {
		idList[i] = gen.Generate()
	}
	for i := 0; i < maxNumber; i++ {
		if !float64Safe(idList[i]) {
			t.Fatal("Unsafe for IEEE-754 64-bit floating-point numbers")
		}
	}
}

func TestIDGen_MaxSafeTime(t *testing.T) {
	gen, err := NewIDGen(WithNode(0, 3), WithStepBits(9))
	if err != nil {
		t.Fatal(err)
	}

	last := gen.MaxSafeTime(Float64SafeBits)
	t.Log("Max safe time:", last)

	last = last.Add(time.Microsecond * 999)
	var idMin, idMax ID
	idMin, idMax = gen.GenerateRangeInNode(last)
	for id := idMin; id <= idMax; id++ {
		if !float64Safe(id) {
			t.Fatal("Excepted safe result")
		}
	}

	last = last.Add(time.Microsecond)
	idMin, idMax = gen.GenerateRangeInNode(last)
	for id := idMin; id <= idMax; id++ {
		if !float64Safe(id) {
			return
		}
	}
	t.Fatal("Excepted Unsafe result")
}
