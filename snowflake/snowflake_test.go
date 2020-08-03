package snowflake

import "testing"

func TestIDGen_Generate(t *testing.T) {
	Epoch = 88834974657

	gen, err := NewIDGen(WithNode(0))
	if err != nil {
		t.Fatal(err)
	}

	id := gen.Generate()
	fid := float64(id)
	t.Log(id, int64(fid)==int64(id))
}
